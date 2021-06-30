package tron

import (
	"sync"
)

type GameState struct {
	GameTime int
	players  sync.Map
	teams    sync.Map
}

func (g *GameState) AddPlayer(p *Player) {
	g.players.Store(p.Id, p)
}

func (g *GameState) GetPlayer(id string) *Player {
	pi, ok := g.players.Load(id)
	if !ok {
		p := &Player{Id: id}
		g.AddPlayer(p)
		return p
	}

	return pi.(*Player)
}

func (g *GameState) PlayersRange(f func(id string, player *Player) bool) {
	g.players.Range(func(_, value interface{}) bool {
		p := value.(*Player)
		return f(p.Id, p)
	})
}

func (g *GameState) RemovePlayer(id string) {
	g.players.Delete(id)
}

func (g *GameState) AddTeam(t *Team) {
	g.teams.Store(t.Id, t)
}

func (g *GameState) GetTeam(id string) *Team {
	ti, ok := g.teams.Load(id)
	if !ok {
		t := &Team{Id: id, Players: map[string]*Player{}}
		g.AddTeam(t)
		return t
	}

	return ti.(*Team)
}

func (g *GameState) RemoveTeam(id string) {
	g.teams.Delete(id)
}

func (g *GameState) ClearAll() {
	clearMap(&g.players)
	clearMap(&g.teams)
}

func clearMap(m *sync.Map) {
	m.Range(func(key, _ interface{}) bool {
		m.Delete(key)
		return true
	})
}
