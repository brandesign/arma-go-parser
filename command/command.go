package command

import (
	"fmt"
	"strings"
	"time"
)

func Raw(cmd string) error {
	_, err := fmt.Println(cmd)

	return err
}

func Rawf(format string, a ...interface{}) error {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	_, err := fmt.Printf(format, a...)

	return err
}

func Say(msg string) error {
	return Rawf("SAY %s", msg)
}

func Pm(playerId, msg string) error {
	return Rawf("PLAYER_MESSAGE %s %s", playerId, msg)
}

func Kick(playerId, reason string) error {
	return Rawf("KICK %s %s", playerId, reason)
}

func Silence(playerId string) error {
	return Rawf("SILENCE %s", playerId)
}

func Voice(playerId string) error {
	return Rawf("VOICE %s", playerId)
}

func CenterMessage(msg string) error {
	return Rawf("CENTER_MESSAGE %s", msg)
}

func Suspend(playerId string, rounds uint) error {
	return Rawf("SUSPEND %s %d", playerId, rounds)
}

func Kill(playerId string) error {
	return Rawf("KILL %s", playerId)
}

func Respawn(playerId string, posX, posY, dirX, dirY float64, msg bool) error {
	return Rawf("RESPAWN_PLAYER %s %d %s %s %d %d", playerId, msg, posX, posY, dirX, dirY)
}

func ConsoleMessage(msg string) error {
	return Rawf("CONSOLE_MESSAGE %s", msg)
}

func AddScoreTeam(teamId string, score int) error {
	return Rawf("ADD_SCORE_TEAM %s %d", teamId, score)
}

func Include(cfgPath string) error {
	return Rawf("INCLUDE %s", cfgPath)
}

func SInclude(cfgPath string) error {
	return Rawf("SINCLUDE %s", cfgPath)
}

func RInclude(cfgPath string) error {
	return Rawf("RINCLUDE %s", cfgPath)
}

func Logf(format string, a ...interface{}) {
	_ = Rawf(fmt.Sprintf("# %s %s", time.Now(), format), a...)
}
