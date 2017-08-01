package main

import "net/http"
import "github.com/ivanthescientist/tournament_service/handlers"
import "github.com/gorilla/mux"

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/fund", handlers.FundHandler).Methods("GET")
	router.HandleFunc("/take", handlers.TakeHandler).Methods("GET")
	router.HandleFunc("/announceTournament", handlers.AnnounceTournamentHandler).Methods("GET")
	router.HandleFunc("/joinTournament", handlers.JoinTournamentHandler).Methods("GET")
	router.HandleFunc("/resultTournament", handlers.ResultTournamentHandler).Methods("POST")
	router.HandleFunc("/balance", handlers.BalanceHandler).Methods("GET")
	router.HandleFunc("/reset", handlers.ResetHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
