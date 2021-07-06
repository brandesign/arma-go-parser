package tron

import (
	"github.com/brandesign/arma-go-parser/event"
	"github.com/brandesign/arma-go-parser/parser"
	"time"
)

func NewConsumer(state *GameState) *Consumer {
	c := &Consumer{state: state}

	subs := []*parser.Subscription{
		event.CycleDestroyed(c),
		event.CycleCreated(c),
		event.PlayerAiEntered(c),
		event.PlayerAiLeft(c),
		event.TeamColoredName(c),
		event.PlayerColoredName(c),
		event.TeamCreated(c),
		event.PlayerEnteredGrid(c),
		event.OnlinePlayer(c),
		event.OnlineTeam(c),
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

func (c *Consumer) OnCycleDestroyed(evt *event.CycleDestroyedEvent) error {
	p := c.state.GetPlayer(evt.PlayerId)
	p.Alive = false

	return nil
}

func (c *Consumer) OnCycleCreated(evt *event.CycleCreatedEvent) error {
	p := c.state.GetPlayer(evt.PlayerId)
	p.Alive = true

	return nil
}

func (c *Consumer) OnPlayerAiEntered(evt *event.PlayerAiEnteredEvent) error {
	p := c.state.GetPlayer(evt.PlayerId)
	p.ScreenName = evt.ScreenName
	p.IsHuman = false

	return nil
}

func (c *Consumer) OnPlayerAiLeft(evt *event.PlayerAiLeftEvent) error {
	c.state.RemovePlayer(evt.PlayerId)

	return nil
}

func (c *Consumer) OnTeamColoredName(evt *event.TeamColoredNameEvent) error {
	t := c.state.GetTeam(evt.TeamId)
	t.ColoredName = evt.ColoredName

	return nil
}

func (c *Consumer) OnPlayerColoredName(evt *event.PlayerColoredNameEvent) error {
	p := c.state.GetPlayer(evt.PlayerId)
	p.ColoredName = evt.ColoredName

	return nil
}

func (c *Consumer) OnTeamCreated(evt *event.TeamCreatedEvent) error {
	t := c.state.GetTeam(evt.TeamId)
	t.ScreenName = evt.Name

	return nil
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
	p := c.state.GetPlayer(evt.PlayerId)
	p.TeamId = ""

	return nil
}

func (c *Consumer) OnTeamPlayerAdded(evt *event.TeamPlayerAddedEvent) error {
	p := c.state.GetPlayer(evt.PlayerId)
	p.TeamId = evt.TeamId

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
	p.TeamId = evt.TeamId
	p.Color = &Color{
		R: evt.Red,
		G: evt.Green,
		B: evt.Blue,
	}

	return nil
}

func (c *Consumer) GetSubscriptions() []*parser.Subscription {
	return c.subs
}
