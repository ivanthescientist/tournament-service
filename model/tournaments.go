package model

import (
	"log"
	"github.com/ivanthescientist/tournament_service/database"
	"database/sql"
)

func CreateTournament(id string, deposit int64) bool {
	tx, err := database.DB.Begin()
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
	var tournamentDeposit int64
	var perPlayerDeposit int64
	var players []string = append(backerIds, playerId)

	rows, err := tx.Query("SELECT deposit FROM tournaments WHERE id = ?;", tournamentId)

	if !rows.Next() {
		tx.Rollback()
		return false
	}

	rows.Scan(&tournamentDeposit)
	rows.Close()
	perPlayerDeposit = (tournamentDeposit) / int64(len(players))

	// First player gets charged a little bit more and gets rewarded a little more as well.
	var depositRemainder = perPlayerDeposit * int64(len(players)) - tournamentDeposit
	var firstPlayerSum = perPlayerDeposit + depositRemainder

	res, err := tx.Exec("UPDATE players SET balance = balance - ? WHERE id = ? AND balance >= ?;", firstPlayerSum, playerId, firstPlayerSum)
	if getRowsAffected(res, err) != 1 {
		tx.Rollback()
		return false
	}

	for _, backerId := range backerIds {
		res, err = tx.Exec("UPDATE players SET balance = balance - ? WHERE id = ? AND balance > ?;",
			perPlayerDeposit,
			backerId,
			perPlayerDeposit)

		if getRowsAffected(res, err) != 1 {
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
			log.Print(err)
			tx.Rollback()
			return false
		}
	}

	tx.Commit()
	return true
}

func ResultTournament(tournamentId string, winners []map[string]int64) bool {
	//tx, err := database.DB.Begin()
	//
	//
	//
	//tx.Commit()
	return true
}

func getRowsAffected(res sql.Result, err error) int64 {
	if err != nil {
		return -1
	}

	rows, _ := res.RowsAffected()

	return rows
}