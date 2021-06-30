package tron

import (
	"time"
)

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
	ScreenName  string
	Ip          string
	Joined      time.Time
	AccessLevel int
	IsHuman     bool
	Position    Vector2
	Dir         Vector2

	Red   int
	Green int
	Blue  int

	Kills     int
	TeamKills int
	Deaths    int
	Suicides  int

	Property
}
