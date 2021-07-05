package tron

import "regexp"

type Team struct {
	Id          string
	ScreenName  string
	ColoredName string

	Property
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
