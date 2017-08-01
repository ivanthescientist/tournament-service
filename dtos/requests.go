package dtos

import "github.com/ivanthescientist/tournament_service/validators"

type FundRequest struct {
	validators.ValidatedDTO
	PlayerId string
	Points   int64
}

type TakeRequest struct {
	validators.ValidatedDTO
	PlayerId string
	Points   int64
}

type AnnounceTournamentRequest struct {
	validators.ValidatedDTO
	TournamentId string
	Deposit      int64
}

type JoinTournamentRequest struct {
	validators.ValidatedDTO
	TournamentId string
	PlayerId     string
	BackerId     []string
}

type TournamentWinner struct {
	PlayerId string `json:"playerId"`
	Prize    int64  `json:"prize"`
}

type ResultTournamentRequest struct {
	validators.ValidatedDTO         `json:"-"`
	TournamentId string             `json:"tournamentId"`
	Winners      []TournamentWinner `json:"winners"`
}

type PlayerBalanceRequest struct {
	validators.ValidatedDTO
	PlayerId string
}