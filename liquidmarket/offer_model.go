package liquidmarket

import (
	"database/sql"
	"fmt"

	"github.com/satori/go.uuid"
)

type Offer struct {
	MarketID      int       `json:"market_id"`
	MarketMakerID uuid.UUID `json:"market_maker_id"`
	Price         float32   `json:"price"`
	Spread        float32   `json:"spread"`
}

func (o *Offer) submitOffer(db *sql.DB) error {
	statement := fmt.Sprintf("CALL offer_submit(%d, unhex(replace('%s', '-', '')), %f, %f)", o.MarketID, o.MarketMakerID, o.Price, o.Spread)
	_, err := db.Exec(statement)
	return err
}
