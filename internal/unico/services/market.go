package services

import (
	"unico/internal/unico/database"
	"unico/internal/unico/entities"
	"unico/internal/unico/models"
	logger "unico/pkg/log"
)

var log *logger.Logger

func init() {
	log = logger.NewLogger()
}

func SearchMarket(search *models.MarketSearchRequest, db database.Database) ([]entities.Market, error) {
	return db.GetAllMarket(search)
}

func CreateMarket(marketreq *models.MarketCreateRequest, db database.Database) (*entities.Market, error) {
	market := &entities.Market{}
	marketreq.ParseModel(market)
	return market, db.CreateMarket(market)
}

func UpdateMarket(market *entities.Market, marketreq *models.MarketUpdateRequest, db database.Database) error {
	marketreq.ParseModel(market)
	return db.UpdateMarket(market)
}
