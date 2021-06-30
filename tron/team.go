package tron

type Team struct {
	Id         string
	ScreenName string
	Players    map[string]*Player

	Property
}

func (t *Team) AddPlayer(p *Player) {
	if t.Players == nil {
		t.Players = map[string]*Player{}
	}

	t.Players[p.Id] = p
}

func (t *Team) RemovePlayer(id string) {
	if t.Players != nil {
		delete(t.Players, id)
	}
}
