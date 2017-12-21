package liquidmarket

import (
	"database/sql"
	"fmt"

	"github.com/satori/go.uuid"
)

type Address struct {
	ID       uuid.UUID `json:"id"`
	Address  string    `json:"address"`
	GoogleID string    `json:"user_id"`
}

func UpdateAddress(db *sql.DB, id uuid.UUID, address string) error {
	statement := fmt.Sprintf("CALL address_update('%s', unhex(replace('%s', '-', '')))", address, id)
	_, err := db.Exec(statement)
	return err
}
func DeleteAddress(db *sql.DB, id uuid.UUID) error {
	statement := fmt.Sprintf("CALL address_delete(unhex(replace('%s', '-', '')))", id)
	_, err := db.Exec(statement)
	return err
}
func (a *Address) CreateAddress(db *sql.DB) error {
	statement := fmt.Sprintf("CALL address_create('%s', '%s');", a.Address, a.GoogleID)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}
func GetAddresss(db *sql.DB, googleId string) ([]Address, error) {
	statement := fmt.Sprintf("SELECT id, address, user_id FROM addresses WHERE user_id = '%s';", googleId)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	addresses := []Address{}
	for rows.Next() {
		var a Address
		if err := rows.Scan(&a.ID, &a.Address, &a.GoogleID); err != nil {
			return nil, err
		}
		addresses = append(addresses, a)
	}
	return addresses, nil
}
