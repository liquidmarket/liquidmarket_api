package liquidmarket

import (
	"database/sql"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Listing struct {
	ID               int       `json:"id"`
	OrganisationName string    `json:"organisation_name"`
	ShortName        string    `json:"short_name"`
	MarketMakerID    uuid.UUID `json:"market_maker_id"`
	MarketMakerName  string    `json:"market_maker_name"`
	TotalShares      int       `json:"total_shares"`
	Price            float32   `json:"price"`
	Spread           float32   `json:"spread"`
	BuyPrice         float32   `json:"buy_price"`
	SellPrice        float32   `json:"sell_price"`
	Prices           []Price   `json:"prices"`
}

type Price struct {
	MarketMakerID uuid.UUID `json:"market_maker_id"`
	Price         float32   `json:"price"`
	Spread        float32   `json:"spread"`
	OfferedAt     time.Time `json:"offered_at"`
}

func GetListingsWithPrices(db *sql.DB) ([]Listing, error) {
	statement := fmt.Sprintf("SELECT m.id, o.`name`, m.`shortname`, m.market_maker_id, mm.name, m.`total`, off.`price` + off.spread as 'buy', off.`price` - off.spread as 'sell', off.`price`, off.spread  FROM markets AS m JOIN organisations AS o ON m.`organisation_id` = o.id JOIN offers AS off ON m.`latest_offer_id` = off.id JOIN marketmakers AS mm ON mm.id = m.market_maker_id;")
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	listings := []Listing{}
	for rows.Next() {
		var m Listing
		if err := rows.Scan(&m.ID, &m.OrganisationName, &m.ShortName, &m.MarketMakerID, &m.MarketMakerName, &m.TotalShares, &m.BuyPrice, &m.SellPrice, &m.Price, &m.Spread); err != nil {
			return nil, err
		}
		priceQuery := fmt.Sprintf("SELECT marketmaker_id, price, spread, offered_at FROM offers WHERE market_id = %d ORDER BY offered_at DESC LIMIT 30", m.ID)
		////
		priceQueryRows, err := db.Query(priceQuery)
		if err != nil {
			return nil, err
		}
		defer priceQueryRows.Close()
		prices := []Price{}
		for priceQueryRows.Next() {
			var p Price
			var timeByteArray string
			if err := priceQueryRows.Scan(&p.MarketMakerID, &p.Price, &p.Spread, &timeByteArray); err != nil {
				return nil, err
			}
			time, err := time.Parse("2006-01-02 15:04:05", timeByteArray)
			if err != nil {
				return nil, err
			}
			p.OfferedAt = time
			prices = append(prices, p)
		}
		m.Prices = prices
		////
		listings = append(listings, m)
	}
	return listings, nil
}
