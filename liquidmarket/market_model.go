package liquidmarket

import (
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Market struct {
	ID               int       `json:"id"`
	MarketMakerID    uuid.UUID `json:"market_maker_id"`
	OrganisationName string    `json:"organisation_name"`
	ShortName        string    `json:"short_name"`
	TotalShares      int       `json:"total_shares"`
	Price            float32   `json:"price"`
	Spread           float32   `json:"spread"`
}

func GetMarkets(db *sql.DB) ([]Market, error) {
	statement := fmt.Sprintf("SELECT m.id, m.market_maker_id, o.`name`, m.`shortname`, m.`total`, off.`price`, off.spread FROM markets AS m JOIN organisations AS o ON m.`organisation_id` = o.id JOIN offers AS off ON m.`latest_offer_id` = off.id;")
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	markets := []Market{}
	for rows.Next() {
		var m Market
		if err := rows.Scan(&m.ID, &m.MarketMakerID, &m.OrganisationName, &m.ShortName, &m.TotalShares, &m.Price, &m.Spread); err != nil {
			return nil, err
		}
		markets = append(markets, m)
	}
	return markets, nil
}
