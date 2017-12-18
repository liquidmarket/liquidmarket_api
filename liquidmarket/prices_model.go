package liquidmarket

import (
    "database/sql"
    "fmt"
)

type Market struct {
    OrganisationName    string    `json:"organisation_name"`
    ShortName  string `json:"short_name"`
    TotalShares  int `json:"total_shares"`
    BuyPrice   float32    `json:"buy_price"`
    SellPrice   float32    `json:"sell_price"`
}

func GetPrices(db *sql.DB) ([]Market, error) {
    statement := fmt.Sprintf("SELECT o.`name`, m.`shortname`, m.`total`, 0 as 'buy', 0 as 'sell' FROM markets AS m JOIN organisations AS o ON m.`organisation_id` = o.id;")
    rows, err := db.Query(statement)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    markets := []Market{}
    for rows.Next() {
        var m Market
        if err := rows.Scan(&m.OrganisationName, &m.ShortName, &m.TotalShares, &m.BuyPrice, &m.SellPrice); err != nil {
            return nil, err
        }
        markets = append(markets, m)
    }
    return markets, nil
}