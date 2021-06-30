package consumer

import (
	"github.com/brandesign/arma-go-parser/event"
	"github.com/brandesign/arma-go-parser/tron"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameState_OnDeathFrag(t *testing.T) {
	c := getConsumer()
	err := c.OnDeathFrag(&event.DeathFragEvent{
		PreyId:   "prey",
		HunterId: "hunter",
	})
	if err != nil {
		t.Fatal(err)
	}

	prey := c.state.GetPlayer("prey")
	hunter := c.state.GetPlayer("hunter")

	assert.Equal(t, 1, prey.Deaths)
	assert.Equal(t, 1, hunter.Kills)
}

func getConsumer() *GameState {
	return &GameState{state: &tron.GameState{}}
}
