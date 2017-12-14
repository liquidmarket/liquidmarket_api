package liquidmarket

import (
    "database/sql"
    "fmt"
    "github.com/satori/go.uuid"
)
type Account struct {
    ID    uuid.UUID    `json:"id"`
    Name  string `json:"name"`
    Balance   float32    `json:"balance"`
}
func (a *Account) GetAccount(db *sql.DB) error {
    statement := fmt.Sprintf("SELECT uuid, name, balance FROM accounts WHERE uuid = unhex(replace('%s', '-', ''))", a.ID)
    return db.QueryRow(statement).Scan(&a.ID, &a.Name, &a.Balance)
}
func (a *Account) UpdateAccount(db *sql.DB, accountUuid uuid.UUID) error {
    statement := fmt.Sprintf("CALL update_account_name('%s', unhex(replace('%s', '-', '')))", a.Name, accountUuid)
    _, err := db.Exec(statement)
    return err
}
func DeleteAccount(db *sql.DB, accountUuid uuid.UUID) error {    
    statement := fmt.Sprintf("CALL delete_account(unhex(replace('%s', '-', '')))", accountUuid)
    _, err := db.Exec(statement)
    return err
}
func (a *Account) CreateAccount(db *sql.DB, googleId string) error {
    statement := fmt.Sprintf("CALL new_account('%s', '%s');", a.Name, googleId)
    _, err := db.Exec(statement)
    if err != nil {
        return err
    }
    return nil
}
func GetAccounts(db *sql.DB, googleId string) ([]Account, error) {
    statement := fmt.Sprintf("SELECT uuid, name, balance FROM accounts AS a JOIN permissions AS p ON a.uuid = p.account_id WHERE p.user_id = '%s'", googleId)
    rows, err := db.Query(statement)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    accounts := []Account{}
    for rows.Next() {
        var a Account
        var guid []byte
        if err := rows.Scan(&guid, &a.Name, &a.Balance); err != nil {
            return nil, err
        }
        a.ID = uuid.FromBytesOrNil(guid)
        accounts = append(accounts, a)
    }
    return accounts, nil
}