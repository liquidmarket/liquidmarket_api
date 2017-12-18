package liquidmarket

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

type Trade struct {
	OrganisationName string    `json:"organisation_name"`
	ShortName        string    `json:"short_name"`
	Shares           int       `json:"shares_traded"`
	TradeID          uuid.UUID `json:"id"`
	OccuredAt        time.Time `json:"occured_at"`
	TotalPrice       float32   `json:"total_price"`
	BuyerID          uuid.UUID `json:"buyer_id"`
	SellerID         uuid.UUID `json:"seller_id"`
	TradedByID       string    `json:"traded_by_id"`
	BuyerName        string    `json:"buyer_name"`
	SellerName       string    `json:"seller_name"`
	TradedByName     string    `json:"traded_by_name"`
}

func getTrades(db *sql.DB, userID string) ([]Trade, error) {
	statement := fmt.Sprintf("SELECT o.`name`, m.`shortname`, `occured_at`, `shares`, t.id, `total_price`, buy.`uuid`, sell.`uuid`, t.`investor_user_id`, buy.name, sell.name, CONCAT_WS(' ', u.`first_name`, u.`last_name`) AS `name` FROM trades AS t JOIN `markets` AS m ON m.id = t.market_id JOIN `organisations` as o ON o.id = m.`organisation_id` JOIN `accounts` AS sell ON sell.uuid = t.seller_id JOIN `accounts` AS buy ON buy.uuid = t.buyer_id JOIN users as u ON u.`google_id` = t.`investor_user_id` WHERE t.buyer_id in (SELECT `account_id` FROM permissions WHERE user_id = '%s') OR t.`seller_id` in (SELECT `account_id` FROM permissions WHERE user_id = '%s') OR t.investor_user_id = '%s';", userID, userID, userID)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	trades := []Trade{}
	for rows.Next() {
		var t Trade
		var timeByteArray string
		if err := rows.Scan(&t.OrganisationName, &t.ShortName, &timeByteArray, &t.Shares, &t.TradeID, &t.TotalPrice, &t.BuyerID, &t.SellerID, &t.TradedByID, &t.BuyerName, &t.SellerName, &t.TradedByName); err != nil {
			return nil, err
		}
		time, err := time.Parse("2006-01-02 15:04:05", timeByteArray)
		if err != nil {
			return nil, err
		}
		t.OccuredAt = time
		trades = append(trades, t)
	}
	return trades, nil
}
