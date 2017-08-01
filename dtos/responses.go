package dtos

import "github.com/ivanthescientist/tournament_service/validators"

type PlayerBalanceResponse struct {
	validators.ValidatedDTO
	playerId string
	balance  int64
}
