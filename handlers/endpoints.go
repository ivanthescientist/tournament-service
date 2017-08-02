package handlers

import (
	"net/http"
	"fmt"
	"github.com/ivanthescientist/tournament_service/dtos"
	"encoding/json"
	"github.com/ivanthescientist/tournament_service/model"
)

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, "Tournament Application")
}

func FundHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();

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

	if model.FundPlayer(request.PlayerId, request.Points) {
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusConflict)
	}
}

func TakeHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();

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
	
	if model.WithdrawFromPlayer(request.PlayerId, request.Points) {
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusConflict)
	}
}

func AnnounceTournamentHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();

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
	
	if model.CreateTournament(request.TournamentId, request.Deposit) {
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusConflict)
	}
}

func JoinTournamentHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();

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
	
	if model.JoinTournament(request.TournamentId, request.PlayerId, request.BackerId) {
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusConflict)
	}
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
	
	if model.ResultTournament(request.TournamentId, request.Winners) {
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusConflict)
	}
}

func BalanceHandler(response http.ResponseWriter, rawRequest *http.Request) {
	err := rawRequest.ParseForm();

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
	
	responseObject := dtos.PlayerBalanceResponse {
		PlayerId: request.PlayerId,
		Balance: model.GetPlayerBalance(request.PlayerId),
	}
	
	json.NewEncoder(response).Encode(responseObject)
}

func ResetHandler(response http.ResponseWriter, request *http.Request) {
	if model.ResetDatabase() {
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusConflict)
	}
}
