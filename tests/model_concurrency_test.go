package tests

import (
	"testing"
	//"github.com/stretchr/testify/assert"
	//"github.com/ivanthescientist/tournament_service/database"
	//"github.com/ivanthescientist/tournament_service/model"
	//"github.com/ivanthescientist/tournament_service/dtos"
	"time"
	//"fmt"
	"github.com/ivanthescientist/tournament_service/model"
	"github.com/ivanthescientist/tournament_service/database"
	"github.com/stretchr/testify/assert"
	"github.com/ivanthescientist/tournament_service/dtos"
)

var oneMs, _ = time.ParseDuration("1ms")

func TestPlayerBalanceConcurrency(t *testing.T) {
	const THREAD_NUM = 100
	const ITERATION_NUM = 100
	database.Init()
	model.ResetDatabase()
	sync := make(chan interface{})

	for i := 0; i < THREAD_NUM; i++ {
		go func(i int, sync chan interface{}) {
			for j := 0; j < ITERATION_NUM; j++ {
				model.FundPlayer("P1", 5)
				model.WithdrawFromPlayer("P1", 4)
			}
			sync <- nil
		}(i, sync)
	}

	for i := 0; i < THREAD_NUM; i++ {
		<- sync
	}

	assert.Equal(t, int64(THREAD_NUM * ITERATION_NUM), model.GetPlayerBalance("P1"))
}

func TestTournamentConcurrency(t *testing.T) {
	const THREAD_NUM = 10
	const ITERATION_NUM = 10
	database.Init()
	model.ResetDatabase()
	sync := make(chan interface{})

	model.FundPlayer("P1", 100)
	model.FundPlayer("P2", 100)
	model.FundPlayer("P3", 100)
	model.FundPlayer("P4", 100)
	model.CreateTournament("T1", 100)

	for i := 0; i < THREAD_NUM; i++ {
		go func(i int, sync chan interface{}) {
			for j := 0; j < ITERATION_NUM; j++ {

				model.JoinTournament("T1","P1", []string{})
				model.JoinTournament("T1","P2", []string{})
				model.ResultTournament("T1", []dtos.TournamentWinner{
					{
						PlayerId: "P1",
						Prize: 400,
					},
				})
			}
			sync <- nil
		}(i, sync)

		go func(i int, sync chan interface{}) {
			for j := 0; j < ITERATION_NUM; j++ {
				model.JoinTournament("T1","P3", []string{"P4"})
				model.ResultTournament("T1", []dtos.TournamentWinner{
					dtos.TournamentWinner{
						PlayerId: "P1",
						Prize: 400,
					},
				})
			}
			sync <- nil
		}(i, sync)
	}

	for i := 0; i < THREAD_NUM; i++ {
		<- sync
	}

	assert.Equal(t, int64(400), model.GetPlayerBalance("P1"))
	assert.Equal(t, int64(0), model.GetPlayerBalance("P2"))
	assert.Equal(t, int64(50), model.GetPlayerBalance("P3"))
	assert.Equal(t, int64(50), model.GetPlayerBalance("P4"))
}