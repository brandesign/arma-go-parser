package event

import "time"

var (
	Handlers = map[string]func() interface{}{
		"ADMIN_COMMAND":                func() interface{} { return &AdminCommandEvent{} },
		"ADMIN_LOGIN":                  func() interface{} { return &AdminLoginEvent{} },
		"ADMIN_LOGOUT":                 func() interface{} { return &AdminLogoutEvent{} },
		"AUTHORITY_BLURB":              func() interface{} { return &AuthorityBlurbEvent{} },
		"BALL_VANISH":                  func() interface{} { return &BallVanishEvent{} },
		"BASEZONE_CONQUERED":           func() interface{} { return &BasezoneConqueredEvent{} },
		"BASEZONE_CONQUERER":           func() interface{} { return &BasezoneConquererEvent{} },
		"BASEZONE_CONQUERER_TEAM":      func() interface{} { return &BasezoneConquererTeamEvent{} },
		"BASE_ENEMY_RESPAWN":           func() interface{} { return &BaseEnemyRespawnEvent{} },
		"BASE_RESPAWN":                 func() interface{} { return &BaseRespawnEvent{} },
		"DEATH_BLASTZONE_PLAYER_ENTER": func() interface{} { return &DeathBlastzonePlayerEnterEvent{} },
		"CHAT":                         func() interface{} { return &ChatEvent{} },
		"CURRENT_MAP":                  func() interface{} { return &CurrentMapEvent{} },
		"CUSTOM_INVALID_COMMAND":       func() interface{} { return &CustomInvalidCommandEvent{} },
		"CYCLE_CREATED":                func() interface{} { return &CycleCreatedEvent{} },
		"CYCLE_DEATH_TELEPORT":         func() interface{} { return &CycleDeathTeleportEvent{} },
		"CYCLE_DESTROYED":              func() interface{} { return &CycleDestroyedEvent{} },
		"DEATHZONE_ACTIVATED":          func() interface{} { return &DeathzoneActivatedEvent{} },
		"DEATH_BASEZONE_CONQUERED":     func() interface{} { return &DeathBasezoneConqueredEvent{} },
		"DEATH_DEATHSHOT":              func() interface{} { return &DeathDeathshotEvent{} },
		"DEATH_DEATHZONE":              func() interface{} { return &DeathDeathzoneEvent{} },
		"DEATH_DEATHZONE_TEAM":         func() interface{} { return &DeathDeathzoneTeamEvent{} },
		"DEATH_EXPLOSION":              func() interface{} { return &DeathExplosionEvent{} },
		"DEATH_FRAG":                   func() interface{} { return &DeathFragEvent{} },
		"DEATH_RUBBERZONE":             func() interface{} { return &DeathRubberzoneEvent{} },
		"DEATH_SELF_DESTRUCT":          func() interface{} { return &DeathSelfDestructEvent{} },
		"DEATH_SHOT_FRAG":              func() interface{} { return &DeathShotFragEvent{} },
		"DEATH_SHOT_SUICIDE":           func() interface{} { return &DeathShotSuicideEvent{} },
		"DEATH_SHOT_TEAMKILL":          func() interface{} { return &DeathShotTeamkillEvent{} },
		"DEATH_SUICIDE":                func() interface{} { return &DeathSuicideEvent{} },
		"DEATH_TEAMKILL":               func() interface{} { return &DeathTeamkillEvent{} },
		"DEATH_ZOMBIEZONE":             func() interface{} { return &DeathZombiezoneEvent{} },
		"ENCODING":                     func() interface{} { return &EncodingEvent{} },
		"END_CHALLENGE":                func() interface{} { return &EndChallengeEvent{} },
		"FLAG_CONQUEST_ROUND_WIN":      func() interface{} { return &FlagConquestRoundWinEvent{} },
		"FLAG_DROP":                    func() interface{} { return &FlagDropEvent{} },
		"FLAG_HELD":                    func() interface{} { return &FlagHeldEvent{} },
		"FLAG_RETURN":                  func() interface{} { return &FlagReturnEvent{} },
		"FLAG_SCORE":                   func() interface{} { return &FlagScoreEvent{} },
		"FLAG_TAKE":                    func() interface{} { return &FlagTakeEvent{} },
		"FLAG_TIMEOUT":                 func() interface{} { return &FlagTimeoutEvent{} },
		"GAME_END":                     func() interface{} { return &GameEndEvent{} },
		"GAME_TIME":                    func() interface{} { return &GameTimeEvent{} },
		"INVALID_COMMAND":              func() interface{} { return &InvalidCommandEvent{} },
		"MATCH_ENDED":                  func() interface{} { return &MatchEndedEvent{} },
		"MATCH_SCORE":                  func() interface{} { return &MatchScoreEvent{} },
		"MATCH_SCORE_TEAM":             func() interface{} { return &MatchScoreTeamEvent{} },
		"MATCH_WINNER":                 func() interface{} { return &MatchWinnerEvent{} },
		"NEW_MATCH":                    func() interface{} { return &NewMatchEvent{} },
		"NEW_ROUND":                    func() interface{} { return &NewRoundEvent{} },
		"NEW_SET":                      func() interface{} { return &NewSetEvent{} },
		"NEXT_ROUND":                   func() interface{} { return &NextRoundEvent{} },
		"NUM_HUMANS":                   func() interface{} { return &NumHumansEvent{} },
		"OBJECTZONE_PLAYER_ENTERED":    func() interface{} { return &ObjectzonePlayerEnteredEvent{} },
		"OBJECTZONE_PLAYER_LEFT":       func() interface{} { return &ObjectzonePlayerLeftEvent{} },
		"OBJECTZONE_SPAWNED":           func() interface{} { return &ObjectzoneSpawnedEvent{} },
		"OBJECTZONE_ZONE_ENTERED":      func() interface{} { return &ObjectzoneZoneEnteredEvent{} },
		"ONLINE_AI":                    func() interface{} { return &OnlineAiEvent{} },
		"ONLINE_PLAYER":                func() interface{} { return &OnlinePlayerEvent{} },
		"ONLINE_PLAYERS_ALIVE":         func() interface{} { return &OnlinePlayersAliveEvent{} },
		"ONLINE_PLAYERS_COUNT":         func() interface{} { return &OnlinePlayersCountEvent{} },
		"ONLINE_PLAYERS_DEAD":          func() interface{} { return &OnlinePlayersDeadEvent{} },
		"ONLINE_TEAM":                  func() interface{} { return &OnlineTeamEvent{} },
		"PLAYER_AI_ENTERED":            func() interface{} { return &PlayerAiEnteredEvent{} },
		"PLAYER_AI_LEFT":               func() interface{} { return &PlayerAiLeftEvent{} },
		"PLAYER_COLORED_NAME":          func() interface{} { return &PlayerColoredNameEvent{} },
		"PLAYER_ENTERED_GRID":          func() interface{} { return &PlayerEnteredGridEvent{} },
		"PLAYER_ENTERED_SPECTATOR":     func() interface{} { return &PlayerEnteredSpectatorEvent{} },
		"PLAYER_GRIDPOS":               func() interface{} { return &PlayerGridposEvent{} },
		"PLAYER_KILLED":                func() interface{} { return &PlayerKilledEvent{} },
		"PLAYER_LEFT":                  func() interface{} { return &PlayerLeftEvent{} },
		"PLAYER_RENAMED":               func() interface{} { return &PlayerRenamedEvent{} },
		"POSITIONS":                    func() interface{} { return &PositionsEvent{} },
		"QUEUE_FINISHED":               func() interface{} { return &QueueFinishedEvent{} },
		"QUEUE_STARTED":                func() interface{} { return &QueueStartedEvent{} },
		"ROUND_COMMENCING":             func() interface{} { return &RoundCommencingEvent{} },
		"ROUND_ENDED":                  func() interface{} { return &RoundEndedEvent{} },
		"ROUND_FINISHED":               func() interface{} { return &RoundFinishedEvent{} },
		"ROUND_SCORE":                  func() interface{} { return &RoundScoreEvent{} },
		"ROUND_SCORE_TEAM":             func() interface{} { return &RoundScoreTeamEvent{} },
		"ROUND_STARTED":                func() interface{} { return &RoundStartedEvent{} },
		"ROUND_WINNER":                 func() interface{} { return &RoundWinnerEvent{} },
		"SACRIFICE":                    func() interface{} { return &SacrificeEvent{} },
		"SET_WINNER":                   func() interface{} { return &SetWinnerEvent{} },
		"SHUTDOWN":                     func() interface{} { return &ShutdownEvent{} },
		"SOCCER_BALL_PLAYER_ENTERED":   func() interface{} { return &SoccerBallPlayerEnteredEvent{} },
		"SOCCER_GOAL_PLAYER_ENTERED":   func() interface{} { return &SoccerGoalPlayerEnteredEvent{} },
		"SOCCER_GOAL_SCORED":           func() interface{} { return &SoccerGoalScoredEvent{} },
		"SPAWN_POSITION_TEAM":          func() interface{} { return &SpawnPositionTeamEvent{} },
		"START_CHALLENGE":              func() interface{} { return &StartChallengeEvent{} },
		"SVG_CREATED":                  func() interface{} { return &SvgCreatedEvent{} },
		"TACTICAL_POSITION":            func() interface{} { return &TacticalPositionEvent{} },
		"TACTICAL_STATISTICS":          func() interface{} { return &TacticalStatisticsEvent{} },
		"TARGETZONE_CONQUERED":         func() interface{} { return &TargetzoneConqueredEvent{} },
		"TARGETZONE_PLAYER_ENTER":      func() interface{} { return &TargetzonePlayerEnterEvent{} },
		"TARGETZONE_PLAYER_LEFT":       func() interface{} { return &TargetzonePlayerLeftEvent{} },
		"TARGETZONE_TIMEOUT":           func() interface{} { return &TargetzoneTimeoutEvent{} },
		"TEAM_COLORED_NAME":            func() interface{} { return &TeamColoredNameEvent{} },
		"TEAM_CREATED":                 func() interface{} { return &TeamCreatedEvent{} },
		"TEAM_DESTROYED":               func() interface{} { return &TeamDestroyedEvent{} },
		"TEAM_PLAYER_ADDED":            func() interface{} { return &TeamPlayerAddedEvent{} },
		"TEAM_PLAYER_REMOVED":          func() interface{} { return &TeamPlayerRemovedEvent{} },
		"TEAM_RENAMED":                 func() interface{} { return &TeamRenamedEvent{} },
		"VOTER":                        func() interface{} { return &VoterEvent{} },
		"VOTE_CREATED":                 func() interface{} { return &VoteCreatedEvent{} },
		"WAIT_FOR_EXTERNAL_SCRIPT":     func() interface{} { return &WaitForExternalScriptEvent{} },
		"WINZONE_ACTIVATED":            func() interface{} { return &WinzoneActivatedEvent{} },
		"WINZONE_PLAYER_ENTER":         func() interface{} { return &WinzonePlayerEnterEvent{} },
		"ZONE_COLLAPSED":               func() interface{} { return &ZoneCollapsedEvent{} },
		"ZONE_CREATED":                 func() interface{} { return &ZoneCreatedEvent{} },
		"ZONE_GRIDPOS":                 func() interface{} { return &ZoneGridposEvent{} },
		"ZONE_ROUTE_STOPPED":           func() interface{} { return &ZoneRouteStoppedEvent{} },
		"ZONE_SHOT_RELEASED":           func() interface{} { return &ZoneShotReleasedEvent{} },
		"ZONE_SPAWNED":                 func() interface{} { return &ZoneSpawnedEvent{} },
	}
)

// AdminCommandEvent handles ADMIN_COMMAND <name> <setting>
type AdminCommandEvent struct {
	Name    string
	Setting string
}

// AdminLoginEvent handles ADMIN_LOGIN [login_name] [ip_address]
type AdminLoginEvent struct {
	PlayerId string
	Ip       string
}

// AdminLogoutEvent handles ADMIN_LOGOUT [login_name] [ip_address]
type AdminLogoutEvent struct {
	PlayerId string
	Ip       string
}

// AuthorityBlurbEvent handles AUTHORITY_BLURB <blurb> <player> <text>
type AuthorityBlurbEvent struct {
	Blurb    string
	PlayerId string
	Text     string
}

// BallVanishEvent handles BALL_VANISH <object id> <zone_name> <cx> <cy>
type BallVanishEvent struct {
	BallId string
	ZoneId string
	X      float64
	Y      float64
}

// BasezoneConqueredEvent handles BASEZONE_CONQUERED <team> <cx> <cy>
type BasezoneConqueredEvent struct {
	TeamId string
	X      float64
	Y      float64
}

// BasezoneConquererEvent handles BASEZONE_CONQUERER <player> <% of zone>
type BasezoneConquererEvent struct {
	PlayerId   string
	Percentage float64
}

// BasezoneConquererTeamEvent handles BASEZONE_CONQUERER_TEAM <team> <score>
type BasezoneConquererTeamEvent struct {
	TeamId string
	Score  int
}

// BaseEnemyRespawnEvent handles BASE_ENEMY_RESPAWN  <spawner> <spawned>
type BaseEnemyRespawnEvent struct {
	Spawner string
	Spawned string
}

// BaseRespawnEvent handles BASE_RESPAWN <spawner> <spawned>
type BaseRespawnEvent struct {
	Spawner string
	Spawned string
}

// DeathBlastzonePlayerEnterEvent handles DEATH_BLASTZONE_PLAYER_ENTER <player>
type DeathBlastzonePlayerEnterEvent struct {
	PlayerId string
}

// ChatEvent handles CHAT <chatter> [/me] <chat string>
type ChatEvent struct {
	PlayerId string
	Text     string
}

// CurrentMapEvent handles CURRENT_MAP [size_factor] [size_multiplier] [MAP_FILE]
type CurrentMapEvent struct {
	SizeFactor     float64
	SizeMultiplier float64
	File           string
}

// CustomInvalidCommandEvent handles CUSTOM_INVALID_COMMAND <command> <player_log> <ip> <access_level> <params>
type CustomInvalidCommandEvent struct {
	Command     string
	PlayerId    string
	Ip          string
	AccessLevel int
	Params      []string
}

// CycleCreatedEvent handles CYCLE_CREATED [auth_name] [posx] [posy] [dirx] [diry] [team_name] [time]
type CycleCreatedEvent struct {
	PlayerId string
	PosX     float64
	PosY     float64
	DirX     float64
	DirY     float64
	TeamId   string
	Time     int
}

// CycleDeathTeleportEvent handles CYCLE_DEATH_TELEPORT [auth_name] [posx] [posy] [dirx] [diry] [team_name] [time] [reason] [predator]
type CycleDeathTeleportEvent struct {
	PlayerId string
	PosX     float64
	PosY     float64
	DirX     float64
	DirY     float64
	TeamId   string
	Time     int
}

// CycleDestroyedEvent handles CYCLE_DESTROYED [auth_name] [posx] [posy] [dirx] [diry] [team_name] [time] [reason] [predator]
type CycleDestroyedEvent struct {
	PlayerId string
	PosX     float64
	PosY     float64
	DirX     float64
	DirY     float64
	TeamId   string
	Time     int
	Reason   string
	Predator string
}

// DeathzoneActivatedEvent handles DEATHZONE_ACTIVATED [id] [name] [xpos] [ypos]
type DeathzoneActivatedEvent struct {
	ZoneId string
	Name   string
	PosX   float64
	PosY   float64
}

// DeathBasezoneConqueredEvent handles DEATH_BASEZONE_CONQUERED <player> [NO_ENEMIES]
type DeathBasezoneConqueredEvent struct {
	PlayerId  string
	NoEnemies bool
}

// DeathDeathshotEvent handles DEATH_DEATHSHOT <prey> <predator>
type DeathDeathshotEvent struct {
	PreyId   string
	HunterId string
}

// DeathDeathzoneEvent handles DEATH_DEATHZONE <player>
type DeathDeathzoneEvent struct {
	PlayerId string
}

// DeathDeathzoneTeamEvent handles DEATH_DEATHZONE_TEAM <team> <player>
type DeathDeathzoneTeamEvent struct {
	TeamId   string
	PlayerId string
}

// DeathExplosionEvent handles DEATH_EXPLOSION <prey> <predator>
type DeathExplosionEvent struct {
	PreyId   string
	HunterId string
}

// DeathFragEvent handles DEATH_FRAG <prey> <predator>
type DeathFragEvent struct {
	PreyId   string
	HunterId string
}

// DeathRubberzoneEvent handles DEATH_RUBBERZONE <player>
type DeathRubberzoneEvent struct {
	PlayerId string
}

// DeathSelfDestructEvent handles DEATH_SELF_DESTRUCT <prey> <predator>
type DeathSelfDestructEvent struct {
	PreyId   string
	HunterId string
}

// DeathShotFragEvent handles DEATH_SHOT_FRAG <prey> <predator>
type DeathShotFragEvent struct {
	PreyId   string
	HunterId string
}

// DeathShotSuicideEvent handles DEATH_SHOT_SUICIDE <player>
type DeathShotSuicideEvent struct {
	PlayerId string
}

// DeathShotTeamkillEvent handles DEATH_SHOT_TEAMKILL <prey> <predator>
type DeathShotTeamkillEvent struct {
	PreyId   string
	HunterId string
}

// DeathSuicideEvent handles DEATH_SUICIDE <player>
type DeathSuicideEvent struct {
	PlayerId string
}

// DeathTeamkillEvent handles DEATH_TEAMKILL <prey> <predator>
type DeathTeamkillEvent struct {
	PreyId   string
	HunterId string
}

// DeathZombiezoneEvent handles DEATH_ZOMBIEZONE <prey> [predator]
type DeathZombiezoneEvent struct {
	PreyId   string
	HunterId string
}

// EncodingEvent handles ENCODING <charset>. Specifies the encoding for data in ladderlog.txt.
type EncodingEvent struct {
	Charset string
}

// EndChallengeEvent handles END_CHALLENGE [time]
type EndChallengeEvent struct {
	Time time.Time
}

// FlagConquestRoundWinEvent handles FLAG_CONQUEST_ROUND_WIN <player> <flag team>
type FlagConquestRoundWinEvent struct {
	PlayerId string
	TeamId   string
}

// FlagDropEvent handles FLAG_DROP <player> <flag team>
type FlagDropEvent struct {
	PlayerId string
	TeamId   string
}

// FlagHeldEvent handles FLAG_HELD <player>
type FlagHeldEvent struct {
	PlayerId string
}

// FlagReturnEvent handles FLAG_RETURN <player>
type FlagReturnEvent struct {
	PlayerId string
}

// FlagScoreEvent handles FLAG_SCORE <player> <flag team>
type FlagScoreEvent struct {
	PlayerId string
	TeamId   string
}

// FlagTakeEvent handles FLAG_TAKE <player> <flag team>
type FlagTakeEvent struct {
	PlayerId string
	TeamId   string
}

// FlagTimeoutEvent handles FLAG_TIMEOUT <player> <flag team>
type FlagTimeoutEvent struct {
	PlayerId string
	TeamId   string
}

// GameEndEvent handles GAME_END <date and time>
type GameEndEvent struct {
	Time time.Time
}

// GameTimeEvent handles GAME_TIME <time> (see also: GAME_TIME_INTERVAL)
type GameTimeEvent struct {
	Time    int
	Precise float64
}

// InvalidCommandEvent handles INVALID_COMMAND [command] [player_username] [ip_address] [access_level] [params]
type InvalidCommandEvent struct {
	Command     string
	PlayerId    string
	Ip          string
	AccessLevel int
	Params      []string
}

// MatchEndedEvent handles MATCH_ENDED [time]
type MatchEndedEvent struct {
	Time time.Time
}

// MatchScoreEvent handles MATCH_SCORE [player_score] [player_username] [team_name]
type MatchScoreEvent struct {
	Score    int
	PlayerId string
	TeamId   string
}

// MatchScoreTeamEvent handles MATCH_SCORE_TEAM [team_score] [team_name] [sets_won]
type MatchScoreTeamEvent struct {
	Score   int
	TeamId  string
	SetsWon int
}

// MatchWinnerEvent handles MATCH_WINNER <team> <players>
type MatchWinnerEvent struct {
	TeamId    string
	PlayerIds []string
}

// NewMatchEvent handles NEW_MATCH <date and time>
type NewMatchEvent struct {
	Time time.Time
}

// NewRoundEvent handles NEW_ROUND <date and time>
type NewRoundEvent struct {
	Time time.Time
}

// NewSetEvent handles NEW_SET [current_set] [time]
type NewSetEvent struct {
	CurrentSet string
	Time       string
}

// NextRoundEvent handles NEXT_ROUND [next_round_number] [total_rounds] [map_file] [center_message]
type NextRoundEvent struct {
	RoundNumber   int
	TotalRounds   int
	MapFile       string
	CenterMessage string
}

// NumHumansEvent handles NUM_HUMANS <number of humans>
type NumHumansEvent struct {
	Number int
}

// ObjectzonePlayerEnteredEvent handles OBJECTZONE_PLAYER_ENTERED [zone_id] [zone_name] [zone_pos_x] [zone_pos_y] [player_name] [player_pos_x] [player_pos_y] [player_direction_x] [player_direction_y] [game_time]
type ObjectzonePlayerEnteredEvent struct {
	ZoneId     string
	ZoneName   string
	ZonePosX   float64
	ZonePosY   float64
	PlayerPosX float64
	PlayerPosY float64
	PlayerDirX float64
	PlayerDirY float64
	GameTime   int
}

// ObjectzonePlayerLeftEvent handles OBJECTZONE_PLAYER_LEFT [zone_id] [zone_name] [zone_pos_x] [zone_pos_y] [player_name] [player_pos_x] [player_pos_y] [player_direction_x] [player_direction_y] [game_time]
type ObjectzonePlayerLeftEvent struct {
	// TODO: implement
}

// ObjectzoneSpawnedEvent handles OBJECTZONE_SPAWNED [id] [name] [pos_x] [pos_y] [xdir] [ydir]
type ObjectzoneSpawnedEvent struct {
	// TODO: implement
}

// ObjectzoneZoneEnteredEvent handles OBJECTZONE_ZONE_ENTERED [zone_id] [zone_name] [zone_posx] [zone_posy] [target_id] [target_name] [target_pos_x] [target_pos_y] [target_dir_x] [target_dir_y] [game_time]
type ObjectzoneZoneEnteredEvent struct {
	// TODO: implement
}

// OnlineAiEvent handles ONLINE_AI <name> <team> <score>
type OnlineAiEvent struct {
	// TODO: implement
}

// OnlinePlayerEvent handles ONLINE_PLAYER <name> <id> <r> <g> <b> <access_level> <did_login?> [<ping> <team>]
// ONLINE_PLAYER jip 1 11 12 10 20 114 0.49454 team_blue
type OnlinePlayerEvent struct {
	Name        string
	PlayerId    string
	Red         int
	Green       int
	Blue        int
	AccessLevel int
	LoggedIn    bool
	Ping        float64
	TeamId      string
}

// OnlinePlayersAliveEvent handles ONLINE_PLAYERS_ALIVE <player1> <player2> <player3> ...
type OnlinePlayersAliveEvent struct {
	// TODO: implement
}

// OnlinePlayersCountEvent handles ONLINE_PLAYERS_COUNT <humans> <ais> <humans alive> <ai alive> <humans dead> <ai dead>
type OnlinePlayersCountEvent struct {
	// TODO: implement
}

// OnlinePlayersDeadEvent handles ONLINE_PLAYERS_DEAD <player1> <player2> <player3> ...
type OnlinePlayersDeadEvent struct {
	// TODO: implement
}

// OnlineTeamEvent handles ONLINE_TEAM <name> <screen name>
type OnlineTeamEvent struct {
	TeamId string
	Name   string
}

// PlayerAiEnteredEvent handles PLAYER_AI_ENTERED <name> <screen name>
type PlayerAiEnteredEvent struct {
	// TODO: implement
}

// PlayerAiLeftEvent handles PLAYER_AI_LEFT [ai_name]
type PlayerAiLeftEvent struct {
	// TODO: implement
}

// PlayerColoredNameEvent handles PLAYER_COLORED_NAME [player_useranme] [player_colored_name]
type PlayerColoredNameEvent struct {
	// TODO: implement
}

// PlayerEnteredGridEvent handles PLAYER_ENTERED_GRID <name> <IP> <screen name>
type PlayerEnteredGridEvent struct {
	PlayerId   string
	Ip         string
	ScreenName string
}

// PlayerEnteredSpectatorEvent handles PLAYER_ENTERED_SPECTATOR <name> <IP> <screen name>
type PlayerEnteredSpectatorEvent struct {
	// TODO: implement
}

// PlayerGridposEvent handles PLAYER_GRIDPOS [player_username] [pos_x] [pos_y] [dir_x] [dir_y] [cycle_speed] [player_rubber] [cycle_rubber] [team] [player_braking] [player_brake_reservoir]
type PlayerGridposEvent struct {
	// TODO: implement
}

// PlayerKilledEvent handles PLAYER_KILLED [player_username] [ip_address] [pos_x] [pos_y] [dir_x] [dir_y]
type PlayerKilledEvent struct {
	// TODO: implement
}

// PlayerLeftEvent handles PLAYER_LEFT <name> <IP>
type PlayerLeftEvent struct {
	PlayerId string
	Ip       string
}

// PlayerRenamedEvent handles PLAYER_RENAMED <old name> <new name> <ip> <did_login?> <screen name>
type PlayerRenamedEvent struct {
	OldId      string
	NewId      string
	Ip         string
	LoggedIn   bool
	ScreenName string
}

// PositionsEvent handles POSITIONS <team> <player1 player2 ...>
type PositionsEvent struct {
	TeamId    string
	PlayerIds []string
}

// QueueFinishedEvent handles QUEUE_FINISHED [time]
type QueueFinishedEvent struct {
	// TODO: implement
}

// QueueStartedEvent handles QUEUE_STARTED [time]
type QueueStartedEvent struct {
	// TODO: implement
}

// RoundCommencingEvent handles ROUND_COMMENCING [current_round] [total_rounds]
type RoundCommencingEvent struct {
	// TODO: implement
}

// RoundEndedEvent handles ROUND_ENDED [time]
type RoundEndedEvent struct {
	// TODO: implement
}

// RoundFinishedEvent handles ROUND_FINISHED [time]
type RoundFinishedEvent struct {
	// TODO: implement
}

// RoundScoreEvent handles ROUND_SCORE <score difference> <player> [<team>]
type RoundScoreEvent struct {
	Score    int
	PlayerId string
	TeamId   string
}

// RoundScoreTeamEvent handles ROUND_SCORE_TEAM <score difference> <team>
type RoundScoreTeamEvent struct {
	Score  int
	TeamId string
}

// RoundStartedEvent handles ROUND_STARTED [time]
type RoundStartedEvent struct {
	// TODO: implement
}

// RoundWinnerEvent handles ROUND_WINNER <team> <players>
type RoundWinnerEvent struct {
	TeamId    string
	PlayerIds []string
}

// SacrificeEvent handles SACRIFICE <player who used the hole> <player who created the hole> <player owning the wall the hole was made into>
type SacrificeEvent struct {
	User  string
	Maker string
	Owner string
}

// SetWinnerEvent handles SET_WINNER [team_name]
type SetWinnerEvent struct {
	// TODO: implement
}

// ShutdownEvent handles SHUTDOWN <time> when the server has been shut down using exit/quit commands
type ShutdownEvent struct {
	// TODO: implement
}

// SoccerBallPlayerEnteredEvent handles SOCCER_BALL_PLAYER_ENTERED [player_auth_name] [team]
type SoccerBallPlayerEnteredEvent struct {
	// TODO: implement
}

// SoccerGoalPlayerEnteredEvent handles SOCCER_GOAL_PLAYER_ENTERED [player_auth_name] [player_team] [team owner of the goal]
type SoccerGoalPlayerEnteredEvent struct {
	// TODO: implement
}

// SoccerGoalScoredEvent handles SOCCER_GOAL_SCORED <goal's team> <scored team> <scored player> <time>
type SoccerGoalScoredEvent struct {
	// TODO: implement
}

// SpawnPositionTeamEvent handles SPAWN_POSITION_TEAM [team_name] [new_position]
type SpawnPositionTeamEvent struct {
	TeamId   string
	Position int
}

// StartChallengeEvent handles START_CHALLENGE [time]
type StartChallengeEvent struct {
	// TODO: implement
}

// SvgCreatedEvent handles SVG_CREATED
type SvgCreatedEvent struct {
	// TODO: implement
}

// TacticalPositionEvent handles TACTICAL_POSITION [time] [name] [tact_pos]
type TacticalPositionEvent struct {
	// TODO: implement
}

// TacticalStatisticsEvent handles TACTICAL_STATISTICS [tact_pos] [name] [time] [state] [kills]
type TacticalStatisticsEvent struct {
	// TODO: implement
}

// TargetzoneConqueredEvent handles TARGETZONE_CONQUERED <object_id> <zone_name> <cx> <cy> [<player> [<team>]]
type TargetzoneConqueredEvent struct {
	// TODO: implement
}

// TargetzonePlayerEnterEvent handles TARGETZONE_PLAYER_ENTER <object_id> <zone_name> <cx> <cy> <player> <x> <y> <xdir> <ydir> <time>
type TargetzonePlayerEnterEvent struct {
	// TODO: implement
}

// TargetzonePlayerLeftEvent handles TARGETZONE_PLAYER_LEFT <object_id> <zone_name> <cx> <cy> <player> <x> <y> <xdir> <ydir>
type TargetzonePlayerLeftEvent struct {
	// TODO: implement
}

// TargetzoneTimeoutEvent handles TARGETZONE_TIMEOUT <object_id> <zone_name> <cx> <cy>
type TargetzoneTimeoutEvent struct {
	// TODO: implement
}

// TeamColoredNameEvent handles TEAM_COLORED_NAME [team_name] [team_colored_name]
type TeamColoredNameEvent struct {
	// TODO: implement
}

// TeamCreatedEvent handles TEAM_CREATED <team name>
type TeamCreatedEvent struct {
	TeamId string
	Name   string
}

// TeamDestroyedEvent handles TEAM_DESTROYED <team name>
type TeamDestroyedEvent struct {
	TeamId string
}

// TeamPlayerAddedEvent handles TEAM_PLAYER_ADDED <team name> <player>
type TeamPlayerAddedEvent struct {
	TeamId   string
	PlayerId string
}

// TeamPlayerRemovedEvent handles TEAM_PLAYER_REMOVED <team name> <player>
type TeamPlayerRemovedEvent struct {
	TeamId   string
	PlayerId string
}

// TeamRenamedEvent handles TEAM_RENAMED <old team name> <new team name>
type TeamRenamedEvent struct {
	OldId string
	NewId string
}

// VoterEvent handles VOTER [player_name] [0-against|1-for] [description]
type VoterEvent struct {
	// TODO: implement
}

// VoteCreatedEvent handles VOTE_CREATED [suggestor] [description]
type VoteCreatedEvent struct {
	// TODO: implement
}

// WaitForExternalScriptEvent handles WAIT_FOR_EXTERNAL_SCRIPT (see also: WAIT_FOR_EXTERNAL_SCRIPT and WAIT_FOR_EXTERNAL_SCRIPT_TIMEOUT)
type WaitForExternalScriptEvent struct {
	// TODO: implement
}

// WinzoneActivatedEvent handles WINZONE_ACTIVATED [id] [name] [xpos] [ypos]
type WinzoneActivatedEvent struct {
	// TODO: implement
}

// WinzonePlayerEnterEvent handles WINZONE_PLAYER_ENTER <player> <x> <y> <xdir> <ydir> <time>
type WinzonePlayerEnterEvent struct {
	// TODO: implement
}

// ZoneCollapsedEvent handles ZONE_COLLAPSED <zone_id> <object_id> <cx> <cy>
type ZoneCollapsedEvent struct {
	// TODO: implement
}

// ZoneCreatedEvent handles ZONE_CREATED [effect] [id] [name] [xpos] [ypos] [xdir] [ydir]
type ZoneCreatedEvent struct {
	// TODO: implement
}

// ZoneGridposEvent handles ZONE_GRIDPOS [effect] [id] [name] [radius] [growth] [posx] [posy] [velx] [vely] [r] [g] [b]
type ZoneGridposEvent struct {
	// TODO: implement
}

// ZoneRouteStoppedEvent handles ZONE_ROUTE_STOPPED [effect] [id] [name] [posx] [posy] [velx] [vely]
type ZoneRouteStoppedEvent struct {
	// TODO: implement
}

// ZoneShotReleasedEvent handles ZONE_SHOT_RELEASED [0-shot|1-deathshot] [id] [player_name] [zone_pos_x] [zone_pos_y] [zone_dir_x] [zone_dir_y]
type ZoneShotReleasedEvent struct {
	// TODO: implement
}

// ZoneSpawnedEvent handles ZONE_SPAWNED <zone_effect> <object id> <zone_name> <x> <y> <xdir> <ydir>
type ZoneSpawnedEvent struct {
	// TODO: implement
}
