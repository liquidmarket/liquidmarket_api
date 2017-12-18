package liquidmarket

import (
    "database/sql"
    "fmt"
)

type User struct {
    GoogleID    string    `json:"google_id"`
    FirstName  string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email   string    `json:"email"`
}

func (u *User) UpdateUser(db *sql.DB) error {
    statement := fmt.Sprintf("CALL update_user('%s', '%s', '%s', '%s')", u.GoogleID, u.FirstName, u.LastName, u.Email)
    _, err := db.Exec(statement)
    return err
}