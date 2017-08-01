package main

import "net/http"
import "github.com/ivanthescientist/tournament_service/handlers"
import (
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func main() {
	var err error

	router := mux.NewRouter();
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/fund", handlers.FundHandler).Methods("GET")
	router.HandleFunc("/take", handlers.TakeHandler).Methods("GET")
	router.HandleFunc("/announceTournament", handlers.AnnounceTournamentHandler).Methods("GET")
	router.HandleFunc("/joinTournament", handlers.JoinTournamentHandler).Methods("GET")
	router.HandleFunc("/resultTournament", handlers.ResultTournamentHandler).Methods("POST")
	router.HandleFunc("/balance", handlers.BalanceHandler).Methods("GET")
	router.HandleFunc("/reset", handlers.ResetHandler).Methods("GET")

	DB, err = sql.Open("mysql", "root:1111@tcp(localhost:3306)/tournament_db")

	if err != nil {
		log.Fatal("Error opening DB connection: ", err.Error())
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error testing DB connection", err.Error())
	}

	http.ListenAndServe(":8080", router)
}
