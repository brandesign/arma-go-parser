package onslaught

import (
	"fmt"
	"github.com/brandesign/arma-go-parser/command"
	"github.com/brandesign/arma-go-parser/event"
	"github.com/brandesign/arma-go-parser/parser"
	"github.com/brandesign/arma-go-parser/tron"
	"math/rand"
	"strings"
)

const (
	keyRespawns = "respawns"

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

	if err := c.reset(); err != nil {
		return nil, err
	}

	subs := []*parser.Subscription{
		event.NewMatch(c),
		event.DeathTeamkill(c),
		event.DeathSuicide(c),
		event.DeathFrag(c),
		event.BasezoneConquered(c),
		event.GameTime(c),
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

func (c *Consumer) OnNewMatch(evt *event.NewMatchEvent) error {
	return c.reset()
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
	if c.remainingTime() <= 0 && !c.zoneConquered {
		// if the attacking team didn't conquer the zone within round_time destroy them - losers!
		for _, p := range c.ofTeam.Players {
			if p.Alive {
				if err := p.Kill(); err != nil {
					return err
				}
			}
		}

		return nil
	}

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

		if err := c.printRoleMessages(); err != nil {
			return err
		}
	}

	if err := c.printRemainingTime(); err != nil {
		return err
	}

	return c.checkBonus()
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

func (c *Consumer) reset() error {
	c.zoneConquered = false
	c.zoneSpawned = false

	if err := c.OnSpawnPositionTeam(&event.SpawnPositionTeamEvent{
		TeamId:   teamBlue,
		Position: 0,
	}); err != nil {
		return err
	}
	if err := c.OnSpawnPositionTeam(&event.SpawnPositionTeamEvent{
		TeamId:   teamGold,
		Position: 1,
	}); err != nil {
		return err
	}

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
		return command.ConsoleMessage(fmt.Sprintf("%s 0xffffffdied", p.GetColoredName()))
	} else {
		r := c.o.SpawnRadius
		x := randFloat64(250-r, 250+r)
		y := randFloat64(450-r, 450+r)

		if err := command.Respawn(pId, x, y, 0, -1, true); err != nil {
			return err
		}

		respawns++
		p.SetInt(keyRespawns, respawns)

		msg := fmt.Sprintf("%s 0xffffffhas been respawned. 0x00ff00%d 0xffffffrespawns remaining", p.GetColoredName(), c.o.Respawns-respawns)
		if err := command.ConsoleMessage(msg); err != nil {
			return err
		}
	}

	return nil
}

func (c *Consumer) printRoleMessages() error {
	if err := roleMessage("Defend!", c.defTeam); err != nil {
		return err
	}
	if err := roleMessage("Attack!", c.ofTeam); err != nil {
		return err
	}

	return nil
}

func (c *Consumer) printRemainingTime() error {
	if c.zoneConquered {
		return nil
	}

	remaining := c.remainingTime()
	if remaining <= 0 {
		return nil
	}

	if remaining%60 == 0 || remaining == 30 || remaining == 10 || remaining == 5 {
		var time int
		var unit string
		if remaining >= 60 {
			time = remaining / 60
			unit = "minutes"
		} else {
			time = remaining
			unit = "seconds"
		}

		return command.ConsoleMessage(fmt.Sprintf("%d %s remaining", time, unit))
	}

	return nil
}

func (c *Consumer) checkBonus() error {
	if c.zoneConquered {
		return nil
	}

	if c.remainingTime()%c.o.BonusTime != 0 {
		return nil
	}

	if c.gs.GameTime < c.o.BonusTime {
		return nil
	}

	if err := command.AddScoreTeam(c.defTeam.Id, c.o.BonusScore); err != nil {
		return err
	}

	msg := fmt.Sprintf("%d bonus points awarded to %s", c.o.BonusScore, c.defTeam.GetColoredName())

	return command.ConsoleMessage(msg)
}

func roleMessage(msg string, t *tron.Team) error {
	color := t.GetColor()
	name := t.ScreenName
	stars := fmt.Sprintf("%s%s", color, decorateString("", "*", 40))
	msgs := []string{
		stars,
		fmt.Sprintf("%s%s", color, decorateString(fmt.Sprintf("%s %s", name, msg), "*", 40)),
		stars,
	}

	for _, m := range msgs {
		if err := command.ConsoleMessage(m); err != nil {
			return err
		}
	}

	return nil
}

func decorateString(str, decorator string, length int) string {
	if len(str) >= length {
		return str
	}

	if len(str) == 0 {
		return strings.Repeat(decorator, length)
	}

	if len(str)%2 != 0 {
		str += " "
	}

	dec := strings.Repeat(decorator, (length-len(str)-2)/2)
	return fmt.Sprintf("%s %s %s", dec, str, dec)
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
