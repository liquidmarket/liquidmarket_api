package liquidmarket

import (
    "database/sql"
    "fmt"
)

type User struct {
    ID    string    `json:"id"`
    FirstName  string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email   string    `json:"email"`
}

func GetUsers(db *sql.DB) ([]User, error) {
    statement := fmt.Sprintf("SELECT google_id, first_name, last_name, email FROM users;")
    rows, err := db.Query(statement)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    users := []User{}
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, nil
}