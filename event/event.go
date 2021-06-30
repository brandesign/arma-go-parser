package event

type CommandEvent struct {
	Command  string
	PlayerId string
	Text     string
	//command,player:player,text
}

type PlayerEnteredEvent struct {
	PlayerId string
	Ip       string
	Name     string
	//player:player,ip,screen_name
}
