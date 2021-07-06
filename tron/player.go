package tron

import (
	"fmt"
	"github.com/brandesign/arma-go-parser/command"
	"time"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x, y float64) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

type Player struct {
	Id          string
	TeamId      string
	ScreenName  string
	ColoredName string
	Ip          string
	Joined      time.Time
	AccessLevel int
	IsHuman     bool
	Position    Vector2
	Dir         Vector2
	Alive       bool

	Color *Color

	Property
}

func (p *Player) Kill() error {
	if err := command.Kill(p.Id); err != nil {
		return err
	}

	p.Alive = false

	return nil
}

func (p *Player) GetColoredName() string {
	if p.ColoredName != "" {
		return p.ColoredName
	}

	if p.Color != nil {
		return fmt.Sprintf("%s%s", p.Color.ToHex(), p.ScreenName)
	}

	return p.ScreenName
}

func (c Color) ToHex() string {
	return fmt.Sprintf("0x%02X%02X%02X", c.R, c.G, c.B)
}
