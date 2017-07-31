package main

import "net/http"
import "github.com/ivanthescientist/tournament_service/controller"

func main() {
    http.HandleFunc("/", controller.IndexHandler);
    http.HandleFunc("/fund", controller.FundHandler);
    http.HandleFunc("/take", controller.TakeHandler);
    http.HandleFunc("/announceTournament", controller.AnnounceTournamentHandler);
    http.HandleFunc("/joinTournament", controller.JoinTournamentHandler);
    http.HandleFunc("/resultTournament", controller.ResultTournamentHandler);
    http.HandleFunc("/balance", controller.BalanceHandler);
    http.HandleFunc("/reset", controller.ResetHandler);
    http.ListenAndServe(":8080", nil);
}
