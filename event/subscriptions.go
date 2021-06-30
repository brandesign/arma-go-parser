package event

import (
	"github.com/brandesign/arma-go-parser/parser"
)

type AdminCommandHandler interface {
	OnAdminCommand(evt *AdminCommandEvent) error
}

func AdminCommand(handler AdminCommandHandler) *parser.Subscription {
	return parser.NewSubscription("ADMIN_COMMAND", func(evt interface{}) error {
		return handler.OnAdminCommand(evt.(*AdminCommandEvent))
	})
}

type AdminLoginHandler interface {
	OnAdminLogin(evt *AdminLoginEvent) error
}

func AdminLogin(handler AdminLoginHandler) *parser.Subscription {
	return parser.NewSubscription("ADMIN_LOGIN", func(evt interface{}) error {
		return handler.OnAdminLogin(evt.(*AdminLoginEvent))
	})
}

type AdminLogoutHandler interface {
	OnAdminLogout(evt *AdminLogoutEvent) error
}

func AdminLogout(handler AdminLogoutHandler) *parser.Subscription {
	return parser.NewSubscription("ADMIN_LOGOUT", func(evt interface{}) error {
		return handler.OnAdminLogout(evt.(*AdminLogoutEvent))
	})
}

type AuthorityBlurbHandler interface {
	OnAuthorityBlurb(evt *AuthorityBlurbEvent) error
}

func AuthorityBlurb(handler AuthorityBlurbHandler) *parser.Subscription {
	return parser.NewSubscription("AUTHORITY_BLURB", func(evt interface{}) error {
		return handler.OnAuthorityBlurb(evt.(*AuthorityBlurbEvent))
	})
}

type BallVanishHandler interface {
	OnBallVanish(evt *BallVanishEvent) error
}

func BallVanish(handler BallVanishHandler) *parser.Subscription {
	return parser.NewSubscription("BALL_VANISH", func(evt interface{}) error {
		return handler.OnBallVanish(evt.(*BallVanishEvent))
	})
}

type BasezoneConqueredHandler interface {
	OnBasezoneConquered(evt *BasezoneConqueredEvent) error
}

func BasezoneConquered(handler BasezoneConqueredHandler) *parser.Subscription {
	return parser.NewSubscription("BASEZONE_CONQUERED", func(evt interface{}) error {
		return handler.OnBasezoneConquered(evt.(*BasezoneConqueredEvent))
	})
}

type BasezoneConquererHandler interface {
	OnBasezoneConquerer(evt *BasezoneConquererEvent) error
}

func BasezoneConquerer(handler BasezoneConquererHandler) *parser.Subscription {
	return parser.NewSubscription("BASEZONE_CONQUERER", func(evt interface{}) error {
		return handler.OnBasezoneConquerer(evt.(*BasezoneConquererEvent))
	})
}

type BasezoneConquererTeamHandler interface {
	OnBasezoneConquererTeam(evt *BasezoneConquererTeamEvent) error
}

func BasezoneConquererTeam(handler BasezoneConquererTeamHandler) *parser.Subscription {
	return parser.NewSubscription("BASEZONE_CONQUERER_TEAM", func(evt interface{}) error {
		return handler.OnBasezoneConquererTeam(evt.(*BasezoneConquererTeamEvent))
	})
}

type BaseEnemyRespawnHandler interface {
	OnBaseEnemyRespawn(evt *BaseEnemyRespawnEvent) error
}

func BaseEnemyRespawn(handler BaseEnemyRespawnHandler) *parser.Subscription {
	return parser.NewSubscription("BASE_ENEMY_RESPAWN", func(evt interface{}) error {
		return handler.OnBaseEnemyRespawn(evt.(*BaseEnemyRespawnEvent))
	})
}

type BaseRespawnHandler interface {
	OnBaseRespawn(evt *BaseRespawnEvent) error
}

func BaseRespawn(handler BaseRespawnHandler) *parser.Subscription {
	return parser.NewSubscription("BASE_RESPAWN", func(evt interface{}) error {
		return handler.OnBaseRespawn(evt.(*BaseRespawnEvent))
	})
}

type DeathBlastzonePlayerEnterHandler interface {
	OnDeathBlastzonePlayerEnter(evt *DeathBlastzonePlayerEnterEvent) error
}

func DeathBlastzonePlayerEnter(handler DeathBlastzonePlayerEnterHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_BLASTZONE_PLAYER_ENTER", func(evt interface{}) error {
		return handler.OnDeathBlastzonePlayerEnter(evt.(*DeathBlastzonePlayerEnterEvent))
	})
}

type ChatHandler interface {
	OnChat(evt *ChatEvent) error
}

func Chat(handler ChatHandler) *parser.Subscription {
	return parser.NewSubscription("CHAT", func(evt interface{}) error {
		return handler.OnChat(evt.(*ChatEvent))
	})
}

type CurrentMapHandler interface {
	OnCurrentMap(evt *CurrentMapEvent) error
}

func CurrentMap(handler CurrentMapHandler) *parser.Subscription {
	return parser.NewSubscription("CURRENT_MAP", func(evt interface{}) error {
		return handler.OnCurrentMap(evt.(*CurrentMapEvent))
	})
}

type CustomInvalidCommandHandler interface {
	OnCustomInvalidCommand(evt *CustomInvalidCommandEvent) error
}

func CustomInvalidCommand(handler CustomInvalidCommandHandler) *parser.Subscription {
	return parser.NewSubscription("CUSTOM_INVALID_COMMAND", func(evt interface{}) error {
		return handler.OnCustomInvalidCommand(evt.(*CustomInvalidCommandEvent))
	})
}

type CycleCreatedHandler interface {
	OnCycleCreated(evt *CycleCreatedEvent) error
}

func CycleCreated(handler CycleCreatedHandler) *parser.Subscription {
	return parser.NewSubscription("CYCLE_CREATED", func(evt interface{}) error {
		return handler.OnCycleCreated(evt.(*CycleCreatedEvent))
	})
}

type CycleDeathTeleportHandler interface {
	OnCycleDeathTeleport(evt *CycleDeathTeleportEvent) error
}

func CycleDeathTeleport(handler CycleDeathTeleportHandler) *parser.Subscription {
	return parser.NewSubscription("CYCLE_DEATH_TELEPORT", func(evt interface{}) error {
		return handler.OnCycleDeathTeleport(evt.(*CycleDeathTeleportEvent))
	})
}

type CycleDestroyedHandler interface {
	OnCycleDestroyed(evt *CycleDestroyedEvent) error
}

func CycleDestroyed(handler CycleDestroyedHandler) *parser.Subscription {
	return parser.NewSubscription("CYCLE_DESTROYED", func(evt interface{}) error {
		return handler.OnCycleDestroyed(evt.(*CycleDestroyedEvent))
	})
}

type DeathzoneActivatedHandler interface {
	OnDeathzoneActivated(evt *DeathzoneActivatedEvent) error
}

func DeathzoneActivated(handler DeathzoneActivatedHandler) *parser.Subscription {
	return parser.NewSubscription("DEATHZONE_ACTIVATED", func(evt interface{}) error {
		return handler.OnDeathzoneActivated(evt.(*DeathzoneActivatedEvent))
	})
}

type DeathBasezoneConqueredHandler interface {
	OnDeathBasezoneConquered(evt *DeathBasezoneConqueredEvent) error
}

func DeathBasezoneConquered(handler DeathBasezoneConqueredHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_BASEZONE_CONQUERED", func(evt interface{}) error {
		return handler.OnDeathBasezoneConquered(evt.(*DeathBasezoneConqueredEvent))
	})
}

type DeathDeathshotHandler interface {
	OnDeathDeathshot(evt *DeathDeathshotEvent) error
}

func DeathDeathshot(handler DeathDeathshotHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_DEATHSHOT", func(evt interface{}) error {
		return handler.OnDeathDeathshot(evt.(*DeathDeathshotEvent))
	})
}

type DeathDeathzoneHandler interface {
	OnDeathDeathzone(evt *DeathDeathzoneEvent) error
}

func DeathDeathzone(handler DeathDeathzoneHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_DEATHZONE", func(evt interface{}) error {
		return handler.OnDeathDeathzone(evt.(*DeathDeathzoneEvent))
	})
}

type DeathDeathzoneTeamHandler interface {
	OnDeathDeathzoneTeam(evt *DeathDeathzoneTeamEvent) error
}

func DeathDeathzoneTeam(handler DeathDeathzoneTeamHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_DEATHZONE_TEAM", func(evt interface{}) error {
		return handler.OnDeathDeathzoneTeam(evt.(*DeathDeathzoneTeamEvent))
	})
}

type DeathExplosionHandler interface {
	OnDeathExplosion(evt *DeathExplosionEvent) error
}

func DeathExplosion(handler DeathExplosionHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_EXPLOSION", func(evt interface{}) error {
		return handler.OnDeathExplosion(evt.(*DeathExplosionEvent))
	})
}

type DeathFragHandler interface {
	OnDeathFrag(evt *DeathFragEvent) error
}

func DeathFrag(handler DeathFragHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_FRAG", func(evt interface{}) error {
		return handler.OnDeathFrag(evt.(*DeathFragEvent))
	})
}

type DeathRubberzoneHandler interface {
	OnDeathRubberzone(evt *DeathRubberzoneEvent) error
}

func DeathRubberzone(handler DeathRubberzoneHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_RUBBERZONE", func(evt interface{}) error {
		return handler.OnDeathRubberzone(evt.(*DeathRubberzoneEvent))
	})
}

type DeathSelfDestructHandler interface {
	OnDeathSelfDestruct(evt *DeathSelfDestructEvent) error
}

func DeathSelfDestruct(handler DeathSelfDestructHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_SELF_DESTRUCT", func(evt interface{}) error {
		return handler.OnDeathSelfDestruct(evt.(*DeathSelfDestructEvent))
	})
}

type DeathShotFragHandler interface {
	OnDeathShotFrag(evt *DeathShotFragEvent) error
}

func DeathShotFrag(handler DeathShotFragHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_SHOT_FRAG", func(evt interface{}) error {
		return handler.OnDeathShotFrag(evt.(*DeathShotFragEvent))
	})
}

type DeathShotSuicideHandler interface {
	OnDeathShotSuicide(evt *DeathShotSuicideEvent) error
}

func DeathShotSuicide(handler DeathShotSuicideHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_SHOT_SUICIDE", func(evt interface{}) error {
		return handler.OnDeathShotSuicide(evt.(*DeathShotSuicideEvent))
	})
}

type DeathShotTeamkillHandler interface {
	OnDeathShotTeamkill(evt *DeathShotTeamkillEvent) error
}

func DeathShotTeamkill(handler DeathShotTeamkillHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_SHOT_TEAMKILL", func(evt interface{}) error {
		return handler.OnDeathShotTeamkill(evt.(*DeathShotTeamkillEvent))
	})
}

type DeathSuicideHandler interface {
	OnDeathSuicide(evt *DeathSuicideEvent) error
}

func DeathSuicide(handler DeathSuicideHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_SUICIDE", func(evt interface{}) error {
		return handler.OnDeathSuicide(evt.(*DeathSuicideEvent))
	})
}

type DeathTeamkillHandler interface {
	OnDeathTeamkill(evt *DeathTeamkillEvent) error
}

func DeathTeamkill(handler DeathTeamkillHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_TEAMKILL", func(evt interface{}) error {
		return handler.OnDeathTeamkill(evt.(*DeathTeamkillEvent))
	})
}

type DeathZombiezoneHandler interface {
	OnDeathZombiezone(evt *DeathZombiezoneEvent) error
}

func DeathZombiezone(handler DeathZombiezoneHandler) *parser.Subscription {
	return parser.NewSubscription("DEATH_ZOMBIEZONE", func(evt interface{}) error {
		return handler.OnDeathZombiezone(evt.(*DeathZombiezoneEvent))
	})
}

type EncodingHandler interface {
	OnEncoding(evt *EncodingEvent) error
}

func Encoding(handler EncodingHandler) *parser.Subscription {
	return parser.NewSubscription("ENCODING", func(evt interface{}) error {
		return handler.OnEncoding(evt.(*EncodingEvent))
	})
}

type EndChallengeHandler interface {
	OnEndChallenge(evt *EndChallengeEvent) error
}

func EndChallenge(handler EndChallengeHandler) *parser.Subscription {
	return parser.NewSubscription("END_CHALLENGE", func(evt interface{}) error {
		return handler.OnEndChallenge(evt.(*EndChallengeEvent))
	})
}

type FlagConquestRoundWinHandler interface {
	OnFlagConquestRoundWin(evt *FlagConquestRoundWinEvent) error
}

func FlagConquestRoundWin(handler FlagConquestRoundWinHandler) *parser.Subscription {
	return parser.NewSubscription("FLAG_CONQUEST_ROUND_WIN", func(evt interface{}) error {
		return handler.OnFlagConquestRoundWin(evt.(*FlagConquestRoundWinEvent))
	})
}

type FlagDropHandler interface {
	OnFlagDrop(evt *FlagDropEvent) error
}

func FlagDrop(handler FlagDropHandler) *parser.Subscription {
	return parser.NewSubscription("FLAG_DROP", func(evt interface{}) error {
		return handler.OnFlagDrop(evt.(*FlagDropEvent))
	})
}

type FlagHeldHandler interface {
	OnFlagHeld(evt *FlagHeldEvent) error
}

func FlagHeld(handler FlagHeldHandler) *parser.Subscription {
	return parser.NewSubscription("FLAG_HELD", func(evt interface{}) error {
		return handler.OnFlagHeld(evt.(*FlagHeldEvent))
	})
}

type FlagReturnHandler interface {
	OnFlagReturn(evt *FlagReturnEvent) error
}

func FlagReturn(handler FlagReturnHandler) *parser.Subscription {
	return parser.NewSubscription("FLAG_RETURN", func(evt interface{}) error {
		return handler.OnFlagReturn(evt.(*FlagReturnEvent))
	})
}

type FlagScoreHandler interface {
	OnFlagScore(evt *FlagScoreEvent) error
}

func FlagScore(handler FlagScoreHandler) *parser.Subscription {
	return parser.NewSubscription("FLAG_SCORE", func(evt interface{}) error {
		return handler.OnFlagScore(evt.(*FlagScoreEvent))
	})
}

type FlagTakeHandler interface {
	OnFlagTake(evt *FlagTakeEvent) error
}

func FlagTake(handler FlagTakeHandler) *parser.Subscription {
	return parser.NewSubscription("FLAG_TAKE", func(evt interface{}) error {
		return handler.OnFlagTake(evt.(*FlagTakeEvent))
	})
}

type FlagTimeoutHandler interface {
	OnFlagTimeout(evt *FlagTimeoutEvent) error
}

func FlagTimeout(handler FlagTimeoutHandler) *parser.Subscription {
	return parser.NewSubscription("FLAG_TIMEOUT", func(evt interface{}) error {
		return handler.OnFlagTimeout(evt.(*FlagTimeoutEvent))
	})
}

type GameEndHandler interface {
	OnGameEnd(evt *GameEndEvent) error
}

func GameEnd(handler GameEndHandler) *parser.Subscription {
	return parser.NewSubscription("GAME_END", func(evt interface{}) error {
		return handler.OnGameEnd(evt.(*GameEndEvent))
	})
}

type GameTimeHandler interface {
	OnGameTime(evt *GameTimeEvent) error
}

func GameTime(handler GameTimeHandler) *parser.Subscription {
	return parser.NewSubscription("GAME_TIME", func(evt interface{}) error {
		return handler.OnGameTime(evt.(*GameTimeEvent))
	})
}

type InvalidCommandHandler interface {
	OnInvalidCommand(evt *InvalidCommandEvent) error
}

func InvalidCommand(handler InvalidCommandHandler) *parser.Subscription {
	return parser.NewSubscription("INVALID_COMMAND", func(evt interface{}) error {
		return handler.OnInvalidCommand(evt.(*InvalidCommandEvent))
	})
}

type MatchEndedHandler interface {
	OnMatchEnded(evt *MatchEndedEvent) error
}

func MatchEnded(handler MatchEndedHandler) *parser.Subscription {
	return parser.NewSubscription("MATCH_ENDED", func(evt interface{}) error {
		return handler.OnMatchEnded(evt.(*MatchEndedEvent))
	})
}

type MatchScoreHandler interface {
	OnMatchScore(evt *MatchScoreEvent) error
}

func MatchScore(handler MatchScoreHandler) *parser.Subscription {
	return parser.NewSubscription("MATCH_SCORE", func(evt interface{}) error {
		return handler.OnMatchScore(evt.(*MatchScoreEvent))
	})
}

type MatchScoreTeamHandler interface {
	OnMatchScoreTeam(evt *MatchScoreTeamEvent) error
}

func MatchScoreTeam(handler MatchScoreTeamHandler) *parser.Subscription {
	return parser.NewSubscription("MATCH_SCORE_TEAM", func(evt interface{}) error {
		return handler.OnMatchScoreTeam(evt.(*MatchScoreTeamEvent))
	})
}

type MatchWinnerHandler interface {
	OnMatchWinner(evt *MatchWinnerEvent) error
}

func MatchWinner(handler MatchWinnerHandler) *parser.Subscription {
	return parser.NewSubscription("MATCH_WINNER", func(evt interface{}) error {
		return handler.OnMatchWinner(evt.(*MatchWinnerEvent))
	})
}

type NewMatchHandler interface {
	OnNewMatch(evt *NewMatchEvent) error
}

func NewMatch(handler NewMatchHandler) *parser.Subscription {
	return parser.NewSubscription("NEW_MATCH", func(evt interface{}) error {
		return handler.OnNewMatch(evt.(*NewMatchEvent))
	})
}

type NewRoundHandler interface {
	OnNewRound(evt *NewRoundEvent) error
}

func NewRound(handler NewRoundHandler) *parser.Subscription {
	return parser.NewSubscription("NEW_ROUND", func(evt interface{}) error {
		return handler.OnNewRound(evt.(*NewRoundEvent))
	})
}

type NewSetHandler interface {
	OnNewSet(evt *NewSetEvent) error
}

func NewSet(handler NewSetHandler) *parser.Subscription {
	return parser.NewSubscription("NEW_SET", func(evt interface{}) error {
		return handler.OnNewSet(evt.(*NewSetEvent))
	})
}

type NextRoundHandler interface {
	OnNextRound(evt *NextRoundEvent) error
}

func NextRound(handler NextRoundHandler) *parser.Subscription {
	return parser.NewSubscription("NEXT_ROUND", func(evt interface{}) error {
		return handler.OnNextRound(evt.(*NextRoundEvent))
	})
}

type NumHumansHandler interface {
	OnNumHumans(evt *NumHumansEvent) error
}

func NumHumans(handler NumHumansHandler) *parser.Subscription {
	return parser.NewSubscription("NUM_HUMANS", func(evt interface{}) error {
		return handler.OnNumHumans(evt.(*NumHumansEvent))
	})
}

type ObjectzonePlayerEnteredHandler interface {
	OnObjectzonePlayerEntered(evt *ObjectzonePlayerEnteredEvent) error
}

func ObjectzonePlayerEntered(handler ObjectzonePlayerEnteredHandler) *parser.Subscription {
	return parser.NewSubscription("OBJECTZONE_PLAYER_ENTERED", func(evt interface{}) error {
		return handler.OnObjectzonePlayerEntered(evt.(*ObjectzonePlayerEnteredEvent))
	})
}

type ObjectzonePlayerLeftHandler interface {
	OnObjectzonePlayerLeft(evt *ObjectzonePlayerLeftEvent) error
}

func ObjectzonePlayerLeft(handler ObjectzonePlayerLeftHandler) *parser.Subscription {
	return parser.NewSubscription("OBJECTZONE_PLAYER_LEFT", func(evt interface{}) error {
		return handler.OnObjectzonePlayerLeft(evt.(*ObjectzonePlayerLeftEvent))
	})
}

type ObjectzoneSpawnedHandler interface {
	OnObjectzoneSpawned(evt *ObjectzoneSpawnedEvent) error
}

func ObjectzoneSpawned(handler ObjectzoneSpawnedHandler) *parser.Subscription {
	return parser.NewSubscription("OBJECTZONE_SPAWNED", func(evt interface{}) error {
		return handler.OnObjectzoneSpawned(evt.(*ObjectzoneSpawnedEvent))
	})
}

type ObjectzoneZoneEnteredHandler interface {
	OnObjectzoneZoneEntered(evt *ObjectzoneZoneEnteredEvent) error
}

func ObjectzoneZoneEntered(handler ObjectzoneZoneEnteredHandler) *parser.Subscription {
	return parser.NewSubscription("OBJECTZONE_ZONE_ENTERED", func(evt interface{}) error {
		return handler.OnObjectzoneZoneEntered(evt.(*ObjectzoneZoneEnteredEvent))
	})
}

type OnlineAiHandler interface {
	OnOnlineAi(evt *OnlineAiEvent) error
}

func OnlineAi(handler OnlineAiHandler) *parser.Subscription {
	return parser.NewSubscription("ONLINE_AI", func(evt interface{}) error {
		return handler.OnOnlineAi(evt.(*OnlineAiEvent))
	})
}

type OnlinePlayerHandler interface {
	OnOnlinePlayer(evt *OnlinePlayerEvent) error
}

func OnlinePlayer(handler OnlinePlayerHandler) *parser.Subscription {
	return parser.NewSubscription("ONLINE_PLAYER", func(evt interface{}) error {
		return handler.OnOnlinePlayer(evt.(*OnlinePlayerEvent))
	})
}

type OnlinePlayersAliveHandler interface {
	OnOnlinePlayersAlive(evt *OnlinePlayersAliveEvent) error
}

func OnlinePlayersAlive(handler OnlinePlayersAliveHandler) *parser.Subscription {
	return parser.NewSubscription("ONLINE_PLAYERS_ALIVE", func(evt interface{}) error {
		return handler.OnOnlinePlayersAlive(evt.(*OnlinePlayersAliveEvent))
	})
}

type OnlinePlayersCountHandler interface {
	OnOnlinePlayersCount(evt *OnlinePlayersCountEvent) error
}

func OnlinePlayersCount(handler OnlinePlayersCountHandler) *parser.Subscription {
	return parser.NewSubscription("ONLINE_PLAYERS_COUNT", func(evt interface{}) error {
		return handler.OnOnlinePlayersCount(evt.(*OnlinePlayersCountEvent))
	})
}

type OnlinePlayersDeadHandler interface {
	OnOnlinePlayersDead(evt *OnlinePlayersDeadEvent) error
}

func OnlinePlayersDead(handler OnlinePlayersDeadHandler) *parser.Subscription {
	return parser.NewSubscription("ONLINE_PLAYERS_DEAD", func(evt interface{}) error {
		return handler.OnOnlinePlayersDead(evt.(*OnlinePlayersDeadEvent))
	})
}

type OnlineTeamHandler interface {
	OnOnlineTeam(evt *OnlineTeamEvent) error
}

func OnlineTeam(handler OnlineTeamHandler) *parser.Subscription {
	return parser.NewSubscription("ONLINE_TEAM", func(evt interface{}) error {
		return handler.OnOnlineTeam(evt.(*OnlineTeamEvent))
	})
}

type PlayerAiEnteredHandler interface {
	OnPlayerAiEntered(evt *PlayerAiEnteredEvent) error
}

func PlayerAiEntered(handler PlayerAiEnteredHandler) *parser.Subscription {
	return parser.NewSubscription("PLAYER_AI_ENTERED", func(evt interface{}) error {
		return handler.OnPlayerAiEntered(evt.(*PlayerAiEnteredEvent))
	})
}

type PlayerAiLeftHandler interface {
	OnPlayerAiLeft(evt *PlayerAiLeftEvent) error
}

func PlayerAiLeft(handler PlayerAiLeftHandler) *parser.Subscription {
	return parser.NewSubscription("PLAYER_AI_LEFT", func(evt interface{}) error {
		return handler.OnPlayerAiLeft(evt.(*PlayerAiLeftEvent))
	})
}

type PlayerColoredNameHandler interface {
	OnPlayerColoredName(evt *PlayerColoredNameEvent) error
}

func PlayerColoredName(handler PlayerColoredNameHandler) *parser.Subscription {
	return parser.NewSubscription("PLAYER_COLORED_NAME", func(evt interface{}) error {
		return handler.OnPlayerColoredName(evt.(*PlayerColoredNameEvent))
	})
}

type PlayerEnteredGridHandler interface {
	OnPlayerEnteredGrid(evt *PlayerEnteredGridEvent) error
}

func PlayerEnteredGrid(handler PlayerEnteredGridHandler) *parser.Subscription {
	return parser.NewSubscription("PLAYER_ENTERED_GRID", func(evt interface{}) error {
		return handler.OnPlayerEnteredGrid(evt.(*PlayerEnteredGridEvent))
	})
}

type PlayerEnteredSpectatorHandler interface {
	OnPlayerEnteredSpectator(evt *PlayerEnteredSpectatorEvent) error
}

func PlayerEnteredSpectator(handler PlayerEnteredSpectatorHandler) *parser.Subscription {
	return parser.NewSubscription("PLAYER_ENTERED_SPECTATOR", func(evt interface{}) error {
		return handler.OnPlayerEnteredSpectator(evt.(*PlayerEnteredSpectatorEvent))
	})
}

type PlayerGridposHandler interface {
	OnPlayerGridpos(evt *PlayerGridposEvent) error
}

func PlayerGridpos(handler PlayerGridposHandler) *parser.Subscription {
	return parser.NewSubscription("PLAYER_GRIDPOS", func(evt interface{}) error {
		return handler.OnPlayerGridpos(evt.(*PlayerGridposEvent))
	})
}

type PlayerKilledHandler interface {
	OnPlayerKilled(evt *PlayerKilledEvent) error
}

func PlayerKilled(handler PlayerKilledHandler) *parser.Subscription {
	return parser.NewSubscription("PLAYER_KILLED", func(evt interface{}) error {
		return handler.OnPlayerKilled(evt.(*PlayerKilledEvent))
	})
}

type PlayerLeftHandler interface {
	OnPlayerLeft(evt *PlayerLeftEvent) error
}

func PlayerLeft(handler PlayerLeftHandler) *parser.Subscription {
	return parser.NewSubscription("PLAYER_LEFT", func(evt interface{}) error {
		return handler.OnPlayerLeft(evt.(*PlayerLeftEvent))
	})
}

type PlayerRenamedHandler interface {
	OnPlayerRenamed(evt *PlayerRenamedEvent) error
}

func PlayerRenamed(handler PlayerRenamedHandler) *parser.Subscription {
	return parser.NewSubscription("PLAYER_RENAMED", func(evt interface{}) error {
		return handler.OnPlayerRenamed(evt.(*PlayerRenamedEvent))
	})
}

type PositionsHandler interface {
	OnPositions(evt *PositionsEvent) error
}

func Positions(handler PositionsHandler) *parser.Subscription {
	return parser.NewSubscription("POSITIONS", func(evt interface{}) error {
		return handler.OnPositions(evt.(*PositionsEvent))
	})
}

type QueueFinishedHandler interface {
	OnQueueFinished(evt *QueueFinishedEvent) error
}

func QueueFinished(handler QueueFinishedHandler) *parser.Subscription {
	return parser.NewSubscription("QUEUE_FINISHED", func(evt interface{}) error {
		return handler.OnQueueFinished(evt.(*QueueFinishedEvent))
	})
}

type QueueStartedHandler interface {
	OnQueueStarted(evt *QueueStartedEvent) error
}

func QueueStarted(handler QueueStartedHandler) *parser.Subscription {
	return parser.NewSubscription("QUEUE_STARTED", func(evt interface{}) error {
		return handler.OnQueueStarted(evt.(*QueueStartedEvent))
	})
}

type RoundCommencingHandler interface {
	OnRoundCommencing(evt *RoundCommencingEvent) error
}

func RoundCommencing(handler RoundCommencingHandler) *parser.Subscription {
	return parser.NewSubscription("ROUND_COMMENCING", func(evt interface{}) error {
		return handler.OnRoundCommencing(evt.(*RoundCommencingEvent))
	})
}

type RoundEndedHandler interface {
	OnRoundEnded(evt *RoundEndedEvent) error
}

func RoundEnded(handler RoundEndedHandler) *parser.Subscription {
	return parser.NewSubscription("ROUND_ENDED", func(evt interface{}) error {
		return handler.OnRoundEnded(evt.(*RoundEndedEvent))
	})
}

type RoundFinishedHandler interface {
	OnRoundFinished(evt *RoundFinishedEvent) error
}

func RoundFinished(handler RoundFinishedHandler) *parser.Subscription {
	return parser.NewSubscription("ROUND_FINISHED", func(evt interface{}) error {
		return handler.OnRoundFinished(evt.(*RoundFinishedEvent))
	})
}

type RoundScoreHandler interface {
	OnRoundScore(evt *RoundScoreEvent) error
}

func RoundScore(handler RoundScoreHandler) *parser.Subscription {
	return parser.NewSubscription("ROUND_SCORE", func(evt interface{}) error {
		return handler.OnRoundScore(evt.(*RoundScoreEvent))
	})
}

type RoundScoreTeamHandler interface {
	OnRoundScoreTeam(evt *RoundScoreTeamEvent) error
}

func RoundScoreTeam(handler RoundScoreTeamHandler) *parser.Subscription {
	return parser.NewSubscription("ROUND_SCORE_TEAM", func(evt interface{}) error {
		return handler.OnRoundScoreTeam(evt.(*RoundScoreTeamEvent))
	})
}

type RoundStartedHandler interface {
	OnRoundStarted(evt *RoundStartedEvent) error
}

func RoundStarted(handler RoundStartedHandler) *parser.Subscription {
	return parser.NewSubscription("ROUND_STARTED", func(evt interface{}) error {
		return handler.OnRoundStarted(evt.(*RoundStartedEvent))
	})
}

type RoundWinnerHandler interface {
	OnRoundWinner(evt *RoundWinnerEvent) error
}

func RoundWinner(handler RoundWinnerHandler) *parser.Subscription {
	return parser.NewSubscription("ROUND_WINNER", func(evt interface{}) error {
		return handler.OnRoundWinner(evt.(*RoundWinnerEvent))
	})
}

type SacrificeHandler interface {
	OnSacrifice(evt *SacrificeEvent) error
}

func Sacrifice(handler SacrificeHandler) *parser.Subscription {
	return parser.NewSubscription("SACRIFICE", func(evt interface{}) error {
		return handler.OnSacrifice(evt.(*SacrificeEvent))
	})
}

type SetWinnerHandler interface {
	OnSetWinner(evt *SetWinnerEvent) error
}

func SetWinner(handler SetWinnerHandler) *parser.Subscription {
	return parser.NewSubscription("SET_WINNER", func(evt interface{}) error {
		return handler.OnSetWinner(evt.(*SetWinnerEvent))
	})
}

type ShutdownHandler interface {
	OnShutdown(evt *ShutdownEvent) error
}

func Shutdown(handler ShutdownHandler) *parser.Subscription {
	return parser.NewSubscription("SHUTDOWN", func(evt interface{}) error {
		return handler.OnShutdown(evt.(*ShutdownEvent))
	})
}

type SoccerBallPlayerEnteredHandler interface {
	OnSoccerBallPlayerEntered(evt *SoccerBallPlayerEnteredEvent) error
}

func SoccerBallPlayerEntered(handler SoccerBallPlayerEnteredHandler) *parser.Subscription {
	return parser.NewSubscription("SOCCER_BALL_PLAYER_ENTERED", func(evt interface{}) error {
		return handler.OnSoccerBallPlayerEntered(evt.(*SoccerBallPlayerEnteredEvent))
	})
}

type SoccerGoalPlayerEnteredHandler interface {
	OnSoccerGoalPlayerEntered(evt *SoccerGoalPlayerEnteredEvent) error
}

func SoccerGoalPlayerEntered(handler SoccerGoalPlayerEnteredHandler) *parser.Subscription {
	return parser.NewSubscription("SOCCER_GOAL_PLAYER_ENTERED", func(evt interface{}) error {
		return handler.OnSoccerGoalPlayerEntered(evt.(*SoccerGoalPlayerEnteredEvent))
	})
}

type SoccerGoalScoredHandler interface {
	OnSoccerGoalScored(evt *SoccerGoalScoredEvent) error
}

func SoccerGoalScored(handler SoccerGoalScoredHandler) *parser.Subscription {
	return parser.NewSubscription("SOCCER_GOAL_SCORED", func(evt interface{}) error {
		return handler.OnSoccerGoalScored(evt.(*SoccerGoalScoredEvent))
	})
}

type SpawnPositionTeamHandler interface {
	OnSpawnPositionTeam(evt *SpawnPositionTeamEvent) error
}

func SpawnPositionTeam(handler SpawnPositionTeamHandler) *parser.Subscription {
	return parser.NewSubscription("SPAWN_POSITION_TEAM", func(evt interface{}) error {
		return handler.OnSpawnPositionTeam(evt.(*SpawnPositionTeamEvent))
	})
}

type StartChallengeHandler interface {
	OnStartChallenge(evt *StartChallengeEvent) error
}

func StartChallenge(handler StartChallengeHandler) *parser.Subscription {
	return parser.NewSubscription("START_CHALLENGE", func(evt interface{}) error {
		return handler.OnStartChallenge(evt.(*StartChallengeEvent))
	})
}

type SvgCreatedHandler interface {
	OnSvgCreated(evt *SvgCreatedEvent) error
}

func SvgCreated(handler SvgCreatedHandler) *parser.Subscription {
	return parser.NewSubscription("SVG_CREATED", func(evt interface{}) error {
		return handler.OnSvgCreated(evt.(*SvgCreatedEvent))
	})
}

type TacticalPositionHandler interface {
	OnTacticalPosition(evt *TacticalPositionEvent) error
}

func TacticalPosition(handler TacticalPositionHandler) *parser.Subscription {
	return parser.NewSubscription("TACTICAL_POSITION", func(evt interface{}) error {
		return handler.OnTacticalPosition(evt.(*TacticalPositionEvent))
	})
}

type TacticalStatisticsHandler interface {
	OnTacticalStatistics(evt *TacticalStatisticsEvent) error
}

func TacticalStatistics(handler TacticalStatisticsHandler) *parser.Subscription {
	return parser.NewSubscription("TACTICAL_STATISTICS", func(evt interface{}) error {
		return handler.OnTacticalStatistics(evt.(*TacticalStatisticsEvent))
	})
}

type TargetzoneConqueredHandler interface {
	OnTargetzoneConquered(evt *TargetzoneConqueredEvent) error
}

func TargetzoneConquered(handler TargetzoneConqueredHandler) *parser.Subscription {
	return parser.NewSubscription("TARGETZONE_CONQUERED", func(evt interface{}) error {
		return handler.OnTargetzoneConquered(evt.(*TargetzoneConqueredEvent))
	})
}

type TargetzonePlayerEnterHandler interface {
	OnTargetzonePlayerEnter(evt *TargetzonePlayerEnterEvent) error
}

func TargetzonePlayerEnter(handler TargetzonePlayerEnterHandler) *parser.Subscription {
	return parser.NewSubscription("TARGETZONE_PLAYER_ENTER", func(evt interface{}) error {
		return handler.OnTargetzonePlayerEnter(evt.(*TargetzonePlayerEnterEvent))
	})
}

type TargetzonePlayerLeftHandler interface {
	OnTargetzonePlayerLeft(evt *TargetzonePlayerLeftEvent) error
}

func TargetzonePlayerLeft(handler TargetzonePlayerLeftHandler) *parser.Subscription {
	return parser.NewSubscription("TARGETZONE_PLAYER_LEFT", func(evt interface{}) error {
		return handler.OnTargetzonePlayerLeft(evt.(*TargetzonePlayerLeftEvent))
	})
}

type TargetzoneTimeoutHandler interface {
	OnTargetzoneTimeout(evt *TargetzoneTimeoutEvent) error
}

func TargetzoneTimeout(handler TargetzoneTimeoutHandler) *parser.Subscription {
	return parser.NewSubscription("TARGETZONE_TIMEOUT", func(evt interface{}) error {
		return handler.OnTargetzoneTimeout(evt.(*TargetzoneTimeoutEvent))
	})
}

type TeamColoredNameHandler interface {
	OnTeamColoredName(evt *TeamColoredNameEvent) error
}

func TeamColoredName(handler TeamColoredNameHandler) *parser.Subscription {
	return parser.NewSubscription("TEAM_COLORED_NAME", func(evt interface{}) error {
		return handler.OnTeamColoredName(evt.(*TeamColoredNameEvent))
	})
}

type TeamCreatedHandler interface {
	OnTeamCreated(evt *TeamCreatedEvent) error
}

func TeamCreated(handler TeamCreatedHandler) *parser.Subscription {
	return parser.NewSubscription("TEAM_CREATED", func(evt interface{}) error {
		return handler.OnTeamCreated(evt.(*TeamCreatedEvent))
	})
}

type TeamDestroyedHandler interface {
	OnTeamDestroyed(evt *TeamDestroyedEvent) error
}

func TeamDestroyed(handler TeamDestroyedHandler) *parser.Subscription {
	return parser.NewSubscription("TEAM_DESTROYED", func(evt interface{}) error {
		return handler.OnTeamDestroyed(evt.(*TeamDestroyedEvent))
	})
}

type TeamPlayerAddedHandler interface {
	OnTeamPlayerAdded(evt *TeamPlayerAddedEvent) error
}

func TeamPlayerAdded(handler TeamPlayerAddedHandler) *parser.Subscription {
	return parser.NewSubscription("TEAM_PLAYER_ADDED", func(evt interface{}) error {
		return handler.OnTeamPlayerAdded(evt.(*TeamPlayerAddedEvent))
	})
}

type TeamPlayerRemovedHandler interface {
	OnTeamPlayerRemoved(evt *TeamPlayerRemovedEvent) error
}

func TeamPlayerRemoved(handler TeamPlayerRemovedHandler) *parser.Subscription {
	return parser.NewSubscription("TEAM_PLAYER_REMOVED", func(evt interface{}) error {
		return handler.OnTeamPlayerRemoved(evt.(*TeamPlayerRemovedEvent))
	})
}

type TeamRenamedHandler interface {
	OnTeamRenamed(evt *TeamRenamedEvent) error
}

func TeamRenamed(handler TeamRenamedHandler) *parser.Subscription {
	return parser.NewSubscription("TEAM_RENAMED", func(evt interface{}) error {
		return handler.OnTeamRenamed(evt.(*TeamRenamedEvent))
	})
}

type VoterHandler interface {
	OnVoter(evt *VoterEvent) error
}

func Voter(handler VoterHandler) *parser.Subscription {
	return parser.NewSubscription("VOTER", func(evt interface{}) error {
		return handler.OnVoter(evt.(*VoterEvent))
	})
}

type VoteCreatedHandler interface {
	OnVoteCreated(evt *VoteCreatedEvent) error
}

func VoteCreated(handler VoteCreatedHandler) *parser.Subscription {
	return parser.NewSubscription("VOTE_CREATED", func(evt interface{}) error {
		return handler.OnVoteCreated(evt.(*VoteCreatedEvent))
	})
}

type WaitForExternalScriptHandler interface {
	OnWaitForExternalScript(evt *WaitForExternalScriptEvent) error
}

func WaitForExternalScript(handler WaitForExternalScriptHandler) *parser.Subscription {
	return parser.NewSubscription("WAIT_FOR_EXTERNAL_SCRIPT", func(evt interface{}) error {
		return handler.OnWaitForExternalScript(evt.(*WaitForExternalScriptEvent))
	})
}

type WinzoneActivatedHandler interface {
	OnWinzoneActivated(evt *WinzoneActivatedEvent) error
}

func WinzoneActivated(handler WinzoneActivatedHandler) *parser.Subscription {
	return parser.NewSubscription("WINZONE_ACTIVATED", func(evt interface{}) error {
		return handler.OnWinzoneActivated(evt.(*WinzoneActivatedEvent))
	})
}

type WinzonePlayerEnterHandler interface {
	OnWinzonePlayerEnter(evt *WinzonePlayerEnterEvent) error
}

func WinzonePlayerEnter(handler WinzonePlayerEnterHandler) *parser.Subscription {
	return parser.NewSubscription("WINZONE_PLAYER_ENTER", func(evt interface{}) error {
		return handler.OnWinzonePlayerEnter(evt.(*WinzonePlayerEnterEvent))
	})
}

type ZoneCollapsedHandler interface {
	OnZoneCollapsed(evt *ZoneCollapsedEvent) error
}

func ZoneCollapsed(handler ZoneCollapsedHandler) *parser.Subscription {
	return parser.NewSubscription("ZONE_COLLAPSED", func(evt interface{}) error {
		return handler.OnZoneCollapsed(evt.(*ZoneCollapsedEvent))
	})
}

type ZoneCreatedHandler interface {
	OnZoneCreated(evt *ZoneCreatedEvent) error
}

func ZoneCreated(handler ZoneCreatedHandler) *parser.Subscription {
	return parser.NewSubscription("ZONE_CREATED", func(evt interface{}) error {
		return handler.OnZoneCreated(evt.(*ZoneCreatedEvent))
	})
}

type ZoneGridposHandler interface {
	OnZoneGridpos(evt *ZoneGridposEvent) error
}

func ZoneGridpos(handler ZoneGridposHandler) *parser.Subscription {
	return parser.NewSubscription("ZONE_GRIDPOS", func(evt interface{}) error {
		return handler.OnZoneGridpos(evt.(*ZoneGridposEvent))
	})
}

type ZoneRouteStoppedHandler interface {
	OnZoneRouteStopped(evt *ZoneRouteStoppedEvent) error
}

func ZoneRouteStopped(handler ZoneRouteStoppedHandler) *parser.Subscription {
	return parser.NewSubscription("ZONE_ROUTE_STOPPED", func(evt interface{}) error {
		return handler.OnZoneRouteStopped(evt.(*ZoneRouteStoppedEvent))
	})
}

type ZoneShotReleasedHandler interface {
	OnZoneShotReleased(evt *ZoneShotReleasedEvent) error
}

func ZoneShotReleased(handler ZoneShotReleasedHandler) *parser.Subscription {
	return parser.NewSubscription("ZONE_SHOT_RELEASED", func(evt interface{}) error {
		return handler.OnZoneShotReleased(evt.(*ZoneShotReleasedEvent))
	})
}

type ZoneSpawnedHandler interface {
	OnZoneSpawned(evt *ZoneSpawnedEvent) error
}

func ZoneSpawned(handler ZoneSpawnedHandler) *parser.Subscription {
	return parser.NewSubscription("ZONE_SPAWNED", func(evt interface{}) error {
		return handler.OnZoneSpawned(evt.(*ZoneSpawnedEvent))
	})
}
