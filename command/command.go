package command

import "fmt"

func Raw(cmd string) error {
	_, err := fmt.Print(cmd)

	return err
}

func Rawf(format string, a ...interface{}) error {
	_, err := fmt.Printf(format, a)

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

func ConsoleMessage(msg string) error {
	return Rawf("CONSOLE_MESSAGE %s", msg)
}

func Comment(msg string) error {
	return Rawf("# %s", msg)
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
