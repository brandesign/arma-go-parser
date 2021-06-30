package tron

type Team struct {
	Id         string
	ScreenName string
	Players    map[string]*Player

	Property
}

func (t *Team) AddPlayer(p *Player) {
	t.Players[p.Id] = p
}

func (t *Team) RemovePlayer(id string) {
	delete(t.Players, id)
}
