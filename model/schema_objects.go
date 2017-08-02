package model

type Player struct {
	id string
	balance int64
}

type Tournament struct {
	id string
	deposit int64
}

type TournamentParticipant struct {
	TournamentId string
	participantId string
	parentId string // references self for original player
}