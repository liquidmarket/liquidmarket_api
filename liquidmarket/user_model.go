package liquidmarket

import (
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	GoogleID  string `json:"google_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (u *User) UpdateUser(db *sql.DB) error {
	statement := fmt.Sprintf("CALL user_update('%s', '%s', '%s', '%s')", u.GoogleID, u.FirstName, u.LastName, u.Email)
	_, err := db.Exec(statement)
	return err
}

func (u *User) GetUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT first_name, last_name, email FROM users WHERE google_id = '%s'", u.GoogleID)
	return db.QueryRow(statement).Scan(&u.FirstName, &u.LastName, &u.Email)
}

func (u *User) getAccountsOrCreate(db *sql.DB) ([]Account, error) {
	statement := fmt.Sprintf("CALL user_create('%s', '%s', '%s', '%s')", u.GoogleID, u.FirstName, u.LastName, u.Email)
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
