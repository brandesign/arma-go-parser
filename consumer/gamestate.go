package consumer

import (
	"github.com/brandesign/arma-go-parser/event"
	"github.com/brandesign/arma-go-parser/parser"
	"github.com/brandesign/arma-go-parser/tron"
	"time"
)

func NewGameState(state *tron.GameState, p *parser.Parser) (*GameState, error) {
	c := &GameState{state: state}

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

	if err := p.Subscribe(subs...); err != nil {
		return nil, err
	}

	return c, nil
}

type GameState struct {
	state *tron.GameState
}

func (g *GameState) OnOnlineTeam(evt *event.OnlineTeamEvent) error {
	t := g.state.GetTeam(evt.TeamId)
	t.ScreenName = evt.Name

	return nil
}

func (g *GameState) OnGameTime(evt *event.GameTimeEvent) error {
	g.state.GameTime = evt.Time

	return nil
}

func (g *GameState) OnPlayerEnteredGrid(evt *event.PlayerEnteredGridEvent) error {
	p := g.state.GetPlayer(evt.PlayerId)

	p.Ip = evt.Ip
	p.ScreenName = evt.ScreenName
	p.Joined = time.Now()
	p.IsHuman = true

	return nil
}

func (g *GameState) OnTeamRenamed(evt *event.TeamRenamedEvent) error {
	t := g.state.GetTeam(evt.OldId)
	g.state.RemoveTeam(evt.OldId)
	t.Id = evt.NewId
	g.state.AddTeam(t)

	return nil
}

func (g *GameState) OnTeamPlayerRemoved(evt *event.TeamPlayerRemovedEvent) error {
	t := g.state.GetTeam(evt.TeamId)
	t.RemovePlayer(evt.PlayerId)

	return nil
}

func (g *GameState) OnTeamPlayerAdded(evt *event.TeamPlayerAddedEvent) error {
	t := g.state.GetTeam(evt.TeamId)
	t.AddPlayer(g.state.GetPlayer(evt.PlayerId))

	return nil
}

func (g *GameState) OnTeamDestroyed(evt *event.TeamDestroyedEvent) error {
	g.state.RemoveTeam(evt.TeamId)

	return nil
}

func (g *GameState) OnPlayerRenamed(evt *event.PlayerRenamedEvent) error {
	p := g.state.GetPlayer(evt.OldId)

	p.Id = evt.NewId
	p.Ip = evt.Ip
	p.ScreenName = evt.ScreenName

	if evt.NewId != evt.OldId {
		g.state.RemovePlayer(evt.OldId)
		g.state.AddPlayer(p)
	}

	return nil
}

func (g *GameState) OnPlayerLeft(evt *event.PlayerLeftEvent) error {
	g.state.RemovePlayer(evt.PlayerId)

	return nil
}

func (g *GameState) OnOnlinePlayer(evt *event.OnlinePlayerEvent) error {
	p := g.state.GetPlayer(evt.PlayerId)

	p.AccessLevel = evt.AccessLevel
	p.Red = evt.Red
	p.Green = evt.Green
	p.Blue = evt.Blue

	if evt.TeamId != "" {
		g.state.GetTeam(evt.TeamId).AddPlayer(p)
	}

	return nil
}

func (g *GameState) OnDeathSuicide(evt *event.DeathSuicideEvent) error {
	p := g.state.GetPlayer(evt.PlayerId)

	p.Deaths++
	p.Suicides++

	return nil
}

func (g *GameState) OnDeathTeamkill(evt *event.DeathTeamkillEvent) error {
	h := g.state.GetPlayer(evt.HunterId)
	p := g.state.GetPlayer(evt.PreyId)

	h.TeamKills++
	p.Deaths++

	return nil
}

func (g *GameState) OnDeathFrag(evt *event.DeathFragEvent) error {
	h := g.state.GetPlayer(evt.HunterId)
	p := g.state.GetPlayer(evt.PreyId)

	h.Kills++
	p.Deaths++

	return nil
}
