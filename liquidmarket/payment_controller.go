package liquidmarket

import (
	"encoding/json"
	"net/http"

	"github.com/satori/go.uuid"
)

func (a *App) deposit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var depositRequest DepositRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&depositRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	account, err := a.Deposit(depositRequest.AccountId, depositRequest.DepositAmount)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error calling SP")
		return
	}
	respondWithJSON(w, http.StatusOK, account)
}

func (a *App) withdrawal(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var withdrawalRequest WithdrawalRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&withdrawalRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	account, err := a.Withdrawal(withdrawalRequest.AccountId, withdrawalRequest.WithdrawalAmount)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error calling SP")
		return
	}
	respondWithJSON(w, http.StatusOK, account)
}

type WithdrawalRequest struct {
	WithdrawalAmount float32   `json:"withdrawal_amount"`
	AccountId        uuid.UUID `json:"account_id"`
}

type DepositRequest struct {
	DepositAmount float32   `json:"deposit_amount"`
	AccountId     uuid.UUID `json:"account_id"`
}
