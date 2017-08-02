package model

import (
	"github.com/ivanthescientist/tournament_service/database"
	"github.com/ivanthescientist/tournament_service/util"
)

func ResetDatabase() bool {
	tx, err := database.DB.Begin()
	util.PrintError("Unable to open transaction for reset", err)

	if err != nil {
		return false;
	}

	_, err = tx.Exec("DELETE FROM tournament_participants;")
	util.PrintError("Unable to reset tournament_participants table", err)
	_, err = tx.Exec("DELETE FROM tournaments;")
	util.PrintError("Unable to reset tournaments table", err)
	_, err = tx.Exec("DELETE FROM players;")
	util.PrintError("Unable to reset players table", err)

	if err != nil {
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true
}
