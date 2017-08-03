package model

import (
	"github.com/ivanthescientist/tournament_service/database"
	"log"
)

func FundPlayer(id string, points int64) bool {
	tx, err := database.DB.Begin()
	stmt, err := tx.Prepare("INSERT INTO `players` (`id`, `balance`) VALUES (?, ?) " +
		"ON DUPLICATE KEY UPDATE `balance` = `balance` + VALUES(`balance`);")
	res, err := stmt.Exec(id, points)
	defer stmt.Close()

	if err == nil && res != nil {
		affected, _ := res.RowsAffected()
		if affected != 0 {
			tx.Commit()
			return true
		}
	} else {
		log.Print("Unsuccessful player fund: ", err.Error())
	}

	tx.Rollback()
	return false;
}

func WithdrawFromPlayer(id string, points int64) bool {
	tx, err := database.DB.Begin()
	stmt, err := tx.Prepare("UPDATE `players` SET `balance` = `balance` - ? WHERE id = ? AND balance > ?;")
	res, err := stmt.Exec(points, id, points)
	defer stmt.Close()

	if err == nil && res != nil {
		affected, _ := res.RowsAffected()
		if affected == 1 {
			tx.Commit()
			return true
		}
	} else {
		log.Print("Unsuccessful player withdrawal: ", err.Error())
	}

	tx.Rollback()
	return false;
}

func GetPlayerBalance(id string) int64 {
	row := database.DB.QueryRow("SELECT `balance` FROM `players` WHERE `id`=?;", id)
	var balance int64
	row.Scan(&balance)
	return balance
}