package onslaught

import (
	"fmt"
	"github.com/brandesign/arma-go-parser/command"
	"github.com/brandesign/arma-go-parser/event"
	"github.com/brandesign/arma-go-parser/parser"
	"github.com/brandesign/arma-go-parser/tron"
	"math/rand"
)

const (
	keyRespawns = "respawns"
	keyColor    = "color"

	teamBlue = "team_blue"
	teamGold = "team_gold"
)

type Options struct {
	RoundTime   int
	BonusTime   int
	BonusScore  int
	Respawns    int
	SpawnRadius float64
}

func NewConsumer(gs *tron.GameState, o Options) (*Consumer, error) {
	c := &Consumer{
		gs: gs,
		o:  o,
	}

	// warm up team setup
	if err := c.OnTeamCreated(&event.TeamCreatedEvent{TeamId: teamBlue}); err != nil {
		return nil, err
	}
	if err := c.OnTeamCreated(&event.TeamCreatedEvent{TeamId: teamGold}); err != nil {
		return nil, err
	}

	if err := c.OnSpawnPositionTeam(&event.SpawnPositionTeamEvent{
		TeamId:   teamBlue,
		Position: 0,
	}); err != nil {
		return nil, err
	}
	if err := c.OnSpawnPositionTeam(&event.SpawnPositionTeamEvent{
		TeamId:   teamGold,
		Position: 1,
	}); err != nil {
		return nil, err
	}

	subs := []*parser.Subscription{
		event.DeathTeamkill(c),
		event.DeathSuicide(c),
		event.DeathFrag(c),
		event.BasezoneConquered(c),
		event.GameTime(c),
		event.TeamCreated(c),
		event.SpawnPositionTeam(c),
		event.NextRound(c),
	}

	c.subs = subs

	return c, nil
}

type Consumer struct {
	gs   *tron.GameState
	subs []*parser.Subscription

	o Options

	zoneSpawned   bool
	zoneConquered bool

	defTeam *tron.Team
	ofTeam  *tron.Team
}

func (c *Consumer) GetSubscriptions() []*parser.Subscription {
	return c.subs
}

func (c *Consumer) OnDeathTeamkill(evt *event.DeathTeamkillEvent) error {
	return c.playerDied(evt.PreyId)
}

func (c *Consumer) OnDeathSuicide(evt *event.DeathSuicideEvent) error {
	return c.playerDied(evt.PlayerId)
}

func (c *Consumer) OnDeathFrag(evt *event.DeathFragEvent) error {
	return c.playerDied(evt.PreyId)
}

func (c *Consumer) OnBasezoneConquered(evt *event.BasezoneConqueredEvent) error {
	c.zoneConquered = true

	return nil
}

func (c *Consumer) OnGameTime(evt *event.GameTimeEvent) error {
	if !c.zoneSpawned {
		/**
		 * spawn zones as early as possible
		 * zones could easily be included in the map, but when spawned from script they are spawned earlier
		 * let helps let people know their role quicker and plan their actions
		 */
		if err := spawnZones(); err != nil {
			return err
		}

		c.zoneSpawned = true

		// TODO: show att def message
	}

	// TODO: $this->timeMessage();
	// TODO: $this->handleBonus();

	if c.remainingTime() <= 0 && !c.zoneConquered {
		// if the attacking team didn't conquer the zone within round_time destroy them - losers!
		for _, p := range c.ofTeam.Players {
			if err := command.Kill(p.Id); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Consumer) OnTeamCreated(evt *event.TeamCreatedEvent) error {
	t := c.gs.GetTeam(evt.TeamId)

	switch t.Id {
	case teamBlue:
		t.SetString(keyColor, "0x4488ff")
		t.ScreenName = "Team Blue"
	case teamGold:
		t.SetString(keyColor, "0xffff44")
		t.ScreenName = "Team Gold"
	default:
		return fmt.Errorf("unknown team = %s", t.Id)
	}

	return nil
}

func (c *Consumer) OnSpawnPositionTeam(evt *event.SpawnPositionTeamEvent) error {
	t := c.gs.GetTeam(evt.TeamId)

	switch evt.Position {
	case 0:
		c.defTeam = t
	case 1:
		c.ofTeam = t
	default:
		return fmt.Errorf("unkown spawn position %d for team %s", evt.Position, t.Id)
	}

	return nil
}

func (c *Consumer) OnNextRound(evt *event.NextRoundEvent) error {
	c.zoneSpawned = false
	c.zoneConquered = false

	c.gs.PlayersRange(func(id string, player *tron.Player) bool {
		player.SetInt(keyRespawns, 0)
		return true
	})

	return nil
}

func (c *Consumer) remainingTime() int {
	return c.o.RoundTime - c.gs.GameTime
}

func (c *Consumer) playerDied(pId string) error {
	p := c.gs.GetPlayer(pId)
	t := c.gs.GetTeamByPlayerId(pId)

	if t == nil {
		return fmt.Errorf("cannot find team for player %s", pId)
	}

	respawns := p.GetInt(keyRespawns)
	if respawns >= c.o.Respawns {
		// TODO: die message
		return nil
	} else {
		r := c.o.SpawnRadius
		x := randFloat64(250-r, 250+r)
		y := randFloat64(450-r, 450+r)

		if err := command.Respawn(pId, x, y, 0, -1, true); err != nil {
			return err
		}

		respawns++
		p.SetInt(keyRespawns, respawns)

		color := t.GetString(keyColor)

		msg := fmt.Sprintf("%s%s 0xffffffhas been respawned. 0x00ff00%d 0xffffffrespawns remaining", color, p.ScreenName, c.o.Respawns-respawns)
		if err := command.ConsoleMessage(msg); err != nil {
			return err
		}
	}

	return nil
}

func spawnZones() error {
	if err := command.Raw("SPAWN_ZONE fortress no_team 250 50 40 0 0 0 false"); err != nil {
		return err
	}

	return command.Raw("SPAWN_ZONE fortress no_team 250 600 1 0 0 0 false")
}

func randFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}