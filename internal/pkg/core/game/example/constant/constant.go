package constant

// Chinese Poker 13 Constants
const (
	PlayerMaxCCHandlerBufferSize                = 256
	RoomMaxCCHandlerBufferSize                  = 1024
	CreditPoolMaxCCHandlerBufferSize            = 256
	StaticDataMaxCCHandlerBufferSize            = 256
	TournamentRoomManagerMaxCCHandlerBufferSize = 1024
	TournamentMaxCCHandlerBufferSize            = 1024

	MinPlayerPerRoom       = 2
	MaxPlayerPerRoom       = 5
	MaxTournament          = 8
	MaxPlayerPerTournament = 512
	ClubMaxRoom            = 256

	CountTickCheckUnderMaintain                 = 60
	CountContinuousTickNoCommandToCloseRoutine  = 600
	DefaultIncreaseBetRuleLevelMinuteSinceBegin = 10
	DefaultIncreaseBetRuleLevelScoreToBet       = 50
	CountTickRoomManagerSyncStore               = 60

	BufferShowCompareResultTimeInSecond = 2
	MinShowCompareResultTimeInSecond    = 30

	TournamentPlayerReconnectOnFailedWaitSecondToRetry                 = 2
	TournamentTickCountToCheckThenAddRoomIfNotEnough                   = 8
	TournamentTickCountToCheckPlayerCountThenCloseRoomToRematchInAsync = 8
	TournamentPrepareToFinalCountDown                                  = 4

	PlayerRoomMatcherBufferSizeOffset            = 4
	PlayerRoomMatcherDelayForPlayerRematchSecond = 1
)
