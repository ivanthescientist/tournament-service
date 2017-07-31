 package controller

import "net/http"
import "fmt"
import "encoding/json"

func IndexHandler(response http.ResponseWriter, request *http.Request)  {
    response.WriteHeader(200);
    fmt.Fprintln(response, "Tournament Application");
}

// GET /fund?playerId={playerId:string}&points={pointsAmount:integer}
func FundHandler(response http.ResponseWriter, request *http.Request)  {
    var playerId = request.URL.Query().Get("playerId");
    var points = request.URL.Query().Get("points");

    response.WriteHeader(200);
    json.NewEncoder(response).Encode("Player #" + playerId + " got " + points + " points.");
}

//
func TakeHandler(response http.ResponseWriter, request *http.Request)  {
    response.WriteHeader(503)
}

func AnnounceTournamentHandler(response http.ResponseWriter, request *http.Request)  {
    response.WriteHeader(503)
}

func JoinTournamentHandler(response http.ResponseWriter, request *http.Request)  {
    response.WriteHeader(503)
}

func ResultTournamentHandler(response http.ResponseWriter, request *http.Request)  {
    response.WriteHeader(503)
}

func BalanceHandler(response http.ResponseWriter, request *http.Request)  {
    response.WriteHeader(503)
}

func ResetHandler(response http.ResponseWriter, request *http.Request)  {
    response.WriteHeader(503)
}