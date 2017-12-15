package liquidmarket

import (
    "database/sql"
    "fmt"
    "github.com/satori/go.uuid"
)

type ShareHolding struct {
    OrganisationName    string    `json:"organisation_name"`
    ShortName  string `json:"short_name"`
	Shares  int `json:"shares"`
    AccountID    uuid.UUID    `json:"account_id"`
    CurrentTradingPrice   float32    `json:"current_price"`
}

func GetShareHoldings(db *sql.DB, account_id uuid.UUID) ([]ShareHolding, error) {
    statement := fmt.Sprintf("SELECT o.`name`, m.`shortname`, sh.`shares`, sh.`account_id`, 0 as 'current_price' FROM shareholdings AS sh JOIN markets AS m ON sh.`market_id` = m.id JOIN organisations AS o ON m.`organisation_id` = o.id WHERE sh.account_id = unhex(replace('%s', '-', ''))", account_id)
    rows, err := db.Query(statement)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    shareholdings := []ShareHolding{}
    for rows.Next() {
        var sh ShareHolding
        if err := rows.Scan(&sh.OrganisationName, &sh.ShortName, &sh.Shares, &sh.AccountID, &sh.CurrentTradingPrice); err != nil {
            return nil, err
        }
        shareholdings = append(shareholdings, sh)
    }
    return shareholdings, nil
}