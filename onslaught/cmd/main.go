package main

import (
	"context"
	"github.com/brandesign/arma-go-parser/event"
	"github.com/brandesign/arma-go-parser/onslaught"
	"github.com/brandesign/arma-go-parser/parser"
	"github.com/brandesign/arma-go-parser/tron"
)

func main() {
	o := onslaught.Options{
		RoundTime:   180,
		BonusTime:   60,
		BonusScore:  5,
		Respawns:    1,
		SpawnRadius: 50,
	}
	gs := &tron.GameState{}
	bc := tron.NewConsumer(gs)
	oc, err := onslaught.NewConsumer(gs, o)
	if err != nil {
		panic(err)
	}

	p, err := parser.NewParser(event.Handlers, bc, oc)
	if err != nil {
		panic(err)
	}

	p.Run(context.Background())
}
