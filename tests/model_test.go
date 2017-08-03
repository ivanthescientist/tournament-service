package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ivanthescientist/tournament_service/database"
	"github.com/ivanthescientist/tournament_service/model"
	"github.com/ivanthescientist/tournament_service/dtos"
)

func TestPlayerBalance(t *testing.T) {
	database.Init()
	model.ResetDatabase()

	assert.Equal(t, int64(0), model.GetPlayerBalance("P1"))
	assert.Equal(t, int64(0), model.GetPlayerBalance("P2"))
	assert.Equal(t, int64(0), model.GetPlayerBalance("P3"))

	model.FundPlayer("P1", 100)
	model.FundPlayer("P2", 200)
	model.FundPlayer("P3", 200)
	model.FundPlayer("P3", 100)
	assert.Equal(t, int64(100), model.GetPlayerBalance("P1"))
	assert.Equal(t, int64(200), model.GetPlayerBalance("P2"))
	assert.Equal(t, int64(300), model.GetPlayerBalance("P3"))

	model.WithdrawFromPlayer("P1", 50)
	assert.Equal(t, int64(50), model.GetPlayerBalance("P1"))

	model.ResetDatabase()
	assert.Equal(t, int64(0), model.GetPlayerBalance("P1"))
	assert.Equal(t, int64(0), model.GetPlayerBalance("P2"))
	assert.Equal(t, int64(0), model.GetPlayerBalance("P3"))
}

func TestTournamentFlow(t *testing.T)  {
	database.Init()
	model.ResetDatabase()

	model.FundPlayer("P1", 100)
	model.FundPlayer("P2", 100)
	model.FundPlayer("P3", 200)
	model.FundPlayer("P4", 200)
	model.FundPlayer("P5", 200)

	assert.Equal(t, int64(100), model.GetPlayerBalance("P1"))
	assert.Equal(t, int64(100), model.GetPlayerBalance("P2"))
	assert.Equal(t, int64(200), model.GetPlayerBalance("P3"))
	assert.Equal(t, int64(200), model.GetPlayerBalance("P4"))
	assert.Equal(t, int64(200), model.GetPlayerBalance("P5"))

	model.CreateTournament("T1", 100)

	model.JoinTournament("T1", "P1", []string{"P2"})
	assert.Equal(t, int64(50), model.GetPlayerBalance("P1"))
	assert.Equal(t, int64(50), model.GetPlayerBalance("P2"))

	model.JoinTournament("T1", "P3", []string{"P4", "P5"})
	assert.Equal(t, int64(166), model.GetPlayerBalance("P3"))
	assert.Equal(t, int64(167), model.GetPlayerBalance("P4"))
	assert.Equal(t, int64(167), model.GetPlayerBalance("P5"))

	// Test that tournament can be joined only once
	model.JoinTournament("T1", "P3", []string{"P4", "P5"})
	assert.Equal(t, int64(166), model.GetPlayerBalance("P3"))
	assert.Equal(t, int64(167), model.GetPlayerBalance("P4"))
	assert.Equal(t, int64(167), model.GetPlayerBalance("P5"))

	model.ResultTournament("T1", []dtos.TournamentWinner{
		dtos.TournamentWinner{
			PlayerId: "P3",
			Prize: 200,
		},
	})

	assert.Equal(t, int64(50), model.GetPlayerBalance("P1"))
	assert.Equal(t, int64(50), model.GetPlayerBalance("P2"))

	assert.Equal(t, int64(234), model.GetPlayerBalance("P3"))
	assert.Equal(t, int64(233), model.GetPlayerBalance("P4"))
	assert.Equal(t, int64(233), model.GetPlayerBalance("P5"))

	model.ResultTournament("T1", []dtos.TournamentWinner{
		dtos.TournamentWinner{
			PlayerId: "P1",
			Prize: 200,
		},
	})

	// Test that tournament is resulted only once.
	assert.Equal(t, int64(50), model.GetPlayerBalance("P1"))
	assert.Equal(t, int64(50), model.GetPlayerBalance("P2"))

	assert.Equal(t, int64(234), model.GetPlayerBalance("P3"))
	assert.Equal(t, int64(233), model.GetPlayerBalance("P4"))
	assert.Equal(t, int64(233), model.GetPlayerBalance("P5"))

	model.ResetDatabase()
}