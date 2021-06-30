package main

import (
	"context"
	"fmt"
	"github.com/brandesign/arma-go-parser/consumer"
	"github.com/brandesign/arma-go-parser/event"
	"github.com/brandesign/arma-go-parser/parser"
	"github.com/brandesign/arma-go-parser/tron"
)

func main() {
	p := parser.NewParser(event.Handlers)
	gs := &tron.GameState{}
	if _, err := consumer.NewGameState(gs, p); err != nil {
		panic(err)
	}

	p.Run(context.Background())
}

type Color struct {
	Red   int
	Green int
	Blue  int
}

func (c Color) ToHexString() string {
	return fmt.Sprintf("0x%x%x%x", c.Red, c.Green, c.Blue)
}

type test struct {
	Str    string
	IntA   int
	FloatA float64
	BoolA  bool
	StrSl  []string
}
