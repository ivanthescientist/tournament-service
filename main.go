package main

import (
	"net/http"
	"github.com/ivanthescientist/tournament_service/handlers"
	"github.com/gorilla/mux"
	"github.com/ivanthescientist/tournament_service/database"
	"github.com/ivanthescientist/tournament_service/model"
	"fmt"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/fund", handlers.FundHandler).Methods("GET")
	router.HandleFunc("/take", handlers.TakeHandler).Methods("GET")
	router.HandleFunc("/announceTournament", handlers.AnnounceTournamentHandler).Methods("GET")
	router.HandleFunc("/joinTournament", handlers.JoinTournamentHandler).Methods("GET")
	router.HandleFunc("/resultTournament", handlers.ResultTournamentHandler).Methods("POST")
	router.HandleFunc("/balance", handlers.BalanceHandler).Methods("GET")
	router.HandleFunc("/reset", handlers.ResetHandler).Methods("GET")

	database.Init()

	model.ResetDatabase()
	model.FundPlayer("P1", 400)
	model.FundPlayer("P2", 400)
	model.FundPlayer("P3", 400)
	model.FundPlayer("P4", 400)

	fmt.Println("P1 - ", model.GetPlayerBalance("P1"))
	fmt.Println("P2 - ", model.GetPlayerBalance("P2"))
	fmt.Println("P3 - ", model.GetPlayerBalance("P3"))
	fmt.Println("P4 - ", model.GetPlayerBalance("P4"))

	model.CreateTournament("T1", 100)
	model.JoinTournament("T1", "P1", []string{})
	model.JoinTournament("T1", "P2", []string{"P3", "P4"})

	fmt.Println("P1 - ", model.GetPlayerBalance("P1"))
	fmt.Println("P2 - ", model.GetPlayerBalance("P2"))
	fmt.Println("P3 - ", model.GetPlayerBalance("P3"))
	fmt.Println("P4 - ", model.GetPlayerBalance("P4"))

	http.ListenAndServe(":8080", router)
}
