package dtos

func (request *FundRequest) IsValid() bool {
	return isValidId(request.PlayerId) && isValidAmount(request.Points)
}

func (request *TakeRequest) IsValid() bool {
	return isValidId(request.PlayerId) && isValidAmount(request.Points)
}

func (request *AnnounceTournamentRequest) IsValid() bool {
	return isValidId(request.TournamentId) && isValidAmount(request.Deposit)
}

func (request *JoinTournamentRequest) IsValid() bool {
	var isValidBackers = true

	if len(request.BackerId) > 0 {
		for i := 0; i < len(request.BackerId); i++ {
			if !isValidId(request.BackerId[i]) {
				isValidBackers = false
			}
		}
	}

	return isValidId(request.TournamentId) && isValidId(request.PlayerId) && isValidBackers
}

func (request *ResultTournamentRequest) IsValid() bool {
	var isValidWinners = true

	for i := 0; i < len(request.Winners); i++ {
		if !isValidId(request.Winners[i].PlayerId) || !isValidAmount(request.Winners[i].Prize) {
			isValidWinners = false
		}
	}

	return isValidId(request.TournamentId) && isValidWinners
}

func (request *PlayerBalanceRequest) IsValid() bool {
	return isValidId(request.PlayerId)
}

func isValidId(id string) bool {
	return len(id) > 0
}

func isValidAmount(amount int64) bool {
	return amount > 0
}
