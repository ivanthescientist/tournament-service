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
	tournament Tournament
	deposit int64 // value deposited, to log transaction
	participant Player
	parent Player // references self for original player
}