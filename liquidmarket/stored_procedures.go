package liquidmarket

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func (a *App) Deposit(accountId uuid.UUID, depositAmount float32) (*Account, error) {
	var account Account
	statement := fmt.Sprintf("CALL deposit(unhex(replace('%s', '-', '')), %f)", accountId, depositAmount)
	err := a.DB.QueryRow(statement).Scan(&account.Name, &account.Balance, &account.ID)
	return &account, err
}

func (a *App) Withdrawal(accountId uuid.UUID, withdrawalAmount float32) (*Account, error) {
	var account Account
	statement := fmt.Sprintf("CALL withdrawal(unhex(replace('%s', '-', '')), %f)", accountId, withdrawalAmount)
	err := a.DB.QueryRow(statement).Scan(&account.Name, &account.Balance, &account.ID)
	return &account, err
}
