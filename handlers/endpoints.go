package handlers

import (
	"net/http"
	"fmt"
	"github.com/ivanthescientist/tournament_service/dtos"
	"encoding/json"
)

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, "Tournament Application")
}

func FundHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

    var request = dtos.FundRequest {
		PlayerId: GetString(rawRequest, "playerId"),
		Points: GetInteger(rawRequest, "points"),
    }


    if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
    }

	// Actual work here
	// Create player with balance - INSERT player ON DUPLICATE KEY UPDATE WHERE

	fmt.Println(queryMap)

    response.WriteHeader(http.StatusOK)
}

func TakeHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var request = dtos.TakeRequest {
		PlayerId: GetString(rawRequest, "playerId"),
		Points: GetInteger(rawRequest, "points"),
	}


	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	// Withdraw amount of player balance - UPDATE WHERE player = id SET balance = MAX(balance - value, 0)
	fmt.Println(queryMap)

	response.WriteHeader(http.StatusOK)
}

func AnnounceTournamentHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var request = dtos.AnnounceTournamentRequest {
		TournamentId: GetString(rawRequest, "tournamentId"),
		Deposit: GetInteger(rawRequest, "deposit"),
	}


	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	// INSERT tournament INTO tournaments - simple
	fmt.Println(queryMap)

	response.WriteHeader(http.StatusOK)
}

func JoinTournamentHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var request = dtos.JoinTournamentRequest {
		PlayerId: GetString(rawRequest, "playerId"),
		BackerId: GetStringArray(rawRequest, "backerId"),
		TournamentId: GetString(rawRequest, "tournamentId"),
	}


	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	
	// Start transaction
	// UPDATE players AS pl LEFT JOIN tournaments AS tr ON pl.id = tr.winner SET balance = balance - amount WHERE
	// pl.balance > amount AND tr.winner IS NULL
	// Check if affected rows == number of players
	// INSERT INTO tournament_records VALUES (... generated set of values ...)
	// Check if affected rows == number of players
	// If rows number match - commit
	// rollback and 401 if not
	fmt.Println(queryMap)

	response.WriteHeader(http.StatusOK)
}

func ResultTournamentHandler(response http.ResponseWriter, rawRequest *http.Request) {
	var request = dtos.ResultTournamentRequest {}

	err := json.NewDecoder(rawRequest.Body).Decode(&request)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	// SELECT * FROM tournament_participants WHERE parent_id = winnerId
	// UPDATE players SET balance + amount
	// check if affected rows == number of players
	// UPDATE tournaments SET winner = playerId WHERE winner IS NULL
	// check if affected rows = 1
	// rollback and 401 if doesn't match
	fmt.Println(request)

	response.WriteHeader(http.StatusOK)
}

func BalanceHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();
	var queryMap = rawRequest.Form;

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var request = dtos.PlayerBalanceRequest {
		PlayerId: GetString(rawRequest, "playerId"),
	}


	if !request.IsValid() {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Actual work here
	// SELECT balance FROM players - simple
	fmt.Println(queryMap)

	response.WriteHeader(http.StatusOK)
}

func ResetHandler(response http.ResponseWriter, request *http.Request) {

	// Actual work here
	// TRUNCATE TABLE tournaments;
	// TRUNCATE TABLE players;
	// TRUNCATE TABLE tournament_participants
	response.WriteHeader(http.StatusOK)
}
