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

type Offer2 struct {
	ID                   uuid.UUID
	MarketID             int
	MarketMakerID        uuid.UUID
	MarketMakerAccountID uuid.UUID
	Price                float32
	Spread               float32
}

func getOffer(db *sql.DB, offerId uuid.UUID) (Offer2, error) {
	var offer Offer2
	statement := fmt.Sprintf("SELECT o.id, market_id, `marketmaker_id`, price, spread, mm.`account_id` FROM offers AS o JOIN marketmakers AS mm on o.`marketmaker_id` = mm.id WHERE o.id = unhex(replace('%s', '-', ''));", offerId)
	err := db.QueryRow(statement).Scan(&offer.ID, &offer.MarketID, &offer.MarketMakerID, &offer.Price, &offer.Spread, &offer.MarketMakerAccountID)
	if err != nil {
		return offer, err
	}
	return offer, nil
}
