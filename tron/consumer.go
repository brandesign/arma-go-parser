package tron

import (
	"github.com/brandesign/arma-go-parser/event"
	"github.com/brandesign/arma-go-parser/parser"
	"time"
)

func NewConsumer(state *GameState) *Consumer {
	c := &Consumer{state: state}

	subs := []*parser.Subscription{
		event.PlayerEnteredGrid(c),
		event.OnlinePlayer(c),
		event.OnlineTeam(c),
		event.DeathSuicide(c),
		event.DeathTeamkill(c),
		event.DeathFrag(c),
		event.PlayerLeft(c),
		event.PlayerRenamed(c),
		event.TeamDestroyed(c),
		event.TeamPlayerAdded(c),
		event.TeamPlayerRemoved(c),
		event.TeamRenamed(c),
		event.GameTime(c),
	}

	c.subs = subs

	return c
}

type Consumer struct {
	state *GameState
	subs  []*parser.Subscription
}

func (c *Consumer) GetSubscriptions() []*parser.Subscription {
	return c.subs
}

func (c *Consumer) OnOnlineTeam(evt *event.OnlineTeamEvent) error {
	t := c.state.GetTeam(evt.TeamId)
	t.ScreenName = evt.Name

	return nil
}

func (c *Consumer) OnGameTime(evt *event.GameTimeEvent) error {
	c.state.GameTime = evt.Time

	return nil
}

func (c *Consumer) OnPlayerEnteredGrid(evt *event.PlayerEnteredGridEvent) error {
	p := c.state.GetPlayer(evt.PlayerId)

	p.Ip = evt.Ip
	p.ScreenName = evt.ScreenName
	p.Joined = time.Now()
	p.IsHuman = true

	return nil
}

func (c *Consumer) OnTeamRenamed(evt *event.TeamRenamedEvent) error {
	t := c.state.GetTeam(evt.OldId)
	c.state.RemoveTeam(evt.OldId)
	t.Id = evt.NewId
	c.state.AddTeam(t)

	return nil
}

func (c *Consumer) OnTeamPlayerRemoved(evt *event.TeamPlayerRemovedEvent) error {
	t := c.state.GetTeam(evt.TeamId)
	t.RemovePlayer(evt.PlayerId)

	return nil
}

func (c *Consumer) OnTeamPlayerAdded(evt *event.TeamPlayerAddedEvent) error {
	t := c.state.GetTeam(evt.TeamId)
	t.AddPlayer(c.state.GetPlayer(evt.PlayerId))

	return nil
}

func (c *Consumer) OnTeamDestroyed(evt *event.TeamDestroyedEvent) error {
	c.state.RemoveTeam(evt.TeamId)

	return nil
}

func (c *Consumer) OnPlayerRenamed(evt *event.PlayerRenamedEvent) error {
	p := c.state.GetPlayer(evt.OldId)

	p.Id = evt.NewId
	p.Ip = evt.Ip
	p.ScreenName = evt.ScreenName

	if evt.NewId != evt.OldId {
		c.state.RemovePlayer(evt.OldId)
		c.state.AddPlayer(p)
	}

	return nil
}

func (c *Consumer) OnPlayerLeft(evt *event.PlayerLeftEvent) error {
	c.state.RemovePlayer(evt.PlayerId)

	return nil
}

func (c *Consumer) OnOnlinePlayer(evt *event.OnlinePlayerEvent) error {
	p := c.state.GetPlayer(evt.PlayerId)

	p.AccessLevel = evt.AccessLevel
	p.Red = evt.Red
	p.Green = evt.Green
	p.Blue = evt.Blue

	t := c.state.GetTeamByPlayerId(p.Id)
	if t != nil && t.Id != evt.TeamId {
		t.RemovePlayer(p.Id)
	}

	if evt.TeamId != "" {
		c.state.GetTeam(evt.TeamId).AddPlayer(p)
	}

	return nil
}

func (c *Consumer) OnDeathSuicide(evt *event.DeathSuicideEvent) error {
	p := c.state.GetPlayer(evt.PlayerId)

	p.Deaths++
	p.Suicides++

	return nil
}

func (c *Consumer) OnDeathTeamkill(evt *event.DeathTeamkillEvent) error {
	h := c.state.GetPlayer(evt.HunterId)
	p := c.state.GetPlayer(evt.PreyId)

	h.TeamKills++
	p.Deaths++

	return nil
}

func (c *Consumer) OnDeathFrag(evt *event.DeathFragEvent) error {
	h := c.state.GetPlayer(evt.HunterId)
	p := c.state.GetPlayer(evt.PreyId)

	h.Kills++
	p.Deaths++

	return nil
}
