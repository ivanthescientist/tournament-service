package model

import (
	"log"
	"github.com/ivanthescientist/tournament_service/database"
	"database/sql"
	"github.com/ivanthescientist/tournament_service/dtos"
)

func CreateTournament(id string, deposit int64) bool {
	tx, err := database.DB.Begin()

	if err != nil {
		log.Println("Failed to open transaction: ", err.Error())
		return false
	}

	stmt, err := tx.Prepare("INSERT INTO `tournaments` (`id`, `deposit`) VALUES (?,?);")
	res, err := stmt.Exec(id, deposit)
	defer stmt.Close()

	if err == nil && res != nil {
		affected, _ := res.RowsAffected()
		if affected != 0 {
			tx.Commit()
			return true
		}
	} else {
		log.Print("Unsuccessful tournament creation: ", err.Error())
	}

	tx.Rollback()
	return false;
}

func JoinTournament(tournamentId string, playerId string, backerIds []string) bool {
	tx, err := database.DB.Begin()

	if err != nil {
		log.Println("Failed to open transaction: ", err.Error())
		return false
	}

	var tournamentDeposit int64
	var perPlayerDeposit int64
	var players []string = append(backerIds, playerId)

	rows, err := tx.Query("SELECT deposit FROM tournaments WHERE id = ?;", tournamentId)

	if !rows.Next() {
		log.Printf("No tournament with ID %s found: %s", tournamentId, err.Error())
		tx.Rollback()
		return false
	}

	rows.Scan(&tournamentDeposit)
	rows.Close()
	perPlayerDeposit = (tournamentDeposit) / int64(len(players))

	// First player gets charged a little bit more and gets rewarded a little more as well.
	var depositRemainder = perPlayerDeposit * int64(len(players)) - tournamentDeposit
	var firstPlayerSum = perPlayerDeposit - depositRemainder

	res, err := tx.Exec("UPDATE players SET balance = balance - ? WHERE id = ? AND balance >= ?;", firstPlayerSum, playerId, firstPlayerSum)
	if getRowsAffected(res, err) != 1 {
		log.Println("Failed to update joining player balance: ", err)
		tx.Rollback()
		return false
	}

	for _, backerId := range backerIds {
		res, err = tx.Exec("UPDATE players SET balance = balance - ? WHERE id = ? AND balance > ?;",
			perPlayerDeposit,
			backerId,
			perPlayerDeposit)

		if getRowsAffected(res, err) != 1 {
			log.Println("Failed to update backing players balances: ", err)
			tx.Rollback()
			return false
		}
	}

	for _, participantId := range players {
		res, err = tx.Exec("INSERT INTO tournament_participants (tournamentId, participantId, parentId) VALUES (?,?,?);",
			tournamentId,
			participantId,
			playerId)

		if getRowsAffected(res, err) != 1 {
			log.Print("Failed to register participants: ", err)
			tx.Rollback()
			return false
		}
	}

	tx.Commit()
	return true
}

func ResultTournament(tournamentId string, winners []dtos.TournamentWinner) bool {
	tx, err := database.DB.Begin()

	if err != nil {
		log.Println("Failed to open transaction: ", err.Error())
		return false
	}

	if len(winners) < 1 {
		log.Println("No winners supplied.")
		return false
	}

	winner := winners[0]

	res, err := tx.Exec("UPDATE tournaments SET winner = ? WHERE id = ? AND winner IS NULL", winner.PlayerId, tournamentId)
	if getRowsAffected(res, err) != 1 {
		log.Print("Failed to update tournament: ", err)
		tx.Rollback()
		return false
	}

	var participants []string
	rows, err := tx.Query("SELECT participantId FROM tournament_participants WHERE parentId = ?", winner.PlayerId)

	if err != nil {
		log.Printf("No participants for tournament %s found: %s", tournamentId, err)
		tx.Rollback()
		return false
	}

	for rows.Next() {
		var participantId string
		rows.Scan(&participantId)
		participants = append(participants, participantId)
	}
	rows.Close()

	var perPlayerWinnings = (winner.Prize) / max(1, int64(len(participants)))
	var winningsRemainder = winner.Prize - perPlayerWinnings * int64(len(participants))

	res, err = tx.Exec(" UPDATE players AS a " +
					"INNER JOIN tournament_participants AS b ON a.id = b.participantId " +
					"SET balance = balance + ? " +
					"WHERE b.parentId = ?;", perPlayerWinnings, winner.PlayerId)

	if getRowsAffected(res, err) != int64(len(participants)) {
		log.Print("Failed to update tournament participant balances: ", err)
		tx.Rollback()
		return false
	}

	if winningsRemainder != 0 {
		res, err = tx.Exec("UPDATE players SET balance = balance + ? WHERE id = ?;", winningsRemainder, winner.PlayerId)
		if getRowsAffected(res, err) != 1 {
			log.Print("Failed to add winnings remainder to main winner's balance: ", err)
			tx.Rollback()
			return false
		}
	}

	tx.Commit()
	return true
}

func getRowsAffected(res sql.Result, err error) int64 {
	if err != nil {
		return -1
	}

	rows, _ := res.RowsAffected()

	return rows
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}