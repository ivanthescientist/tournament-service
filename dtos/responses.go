package dtos

import "github.com/ivanthescientist/tournament_service/validators"

type PlayerBalanceResponse struct {
	validators.ValidatedDTO
	PlayerId string `json:"playerId"`
	Balance  int64  `json:"balance"`
}
