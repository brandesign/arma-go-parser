package tron

import "regexp"

type Team struct {
	Id          string
	ScreenName  string
	ColoredName string
	Players     map[string]*Player

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

func (t *Team) GetColoredName() string {
	if t.ColoredName != "" {
		return t.ColoredName
	}

	return t.ScreenName
}

func (t *Team) GetColor() string {
	n := t.GetColoredName()
	if n == "" {
		return ""
	}

	r, err := regexp.Compile("(?i)0x[0-9a-f]{6}")
	if err != nil {
		return ""
	}
	return r.FindString(n)
}
