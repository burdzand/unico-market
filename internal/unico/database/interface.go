package database

import (
	"unico/internal/unico/entities"
	"unico/internal/unico/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	Database interface {
		GetDB() *gorm.DB
		CreateMarket(market *entities.Market) error
		UpdateMarket(market *entities.Market) error
		GetMarket(id uuid.UUID) (*entities.Market, error)
		GetMarketByRecord(record string) (*entities.Market, error)
		RemoveMarket(record string) error
		GetAllMarket(search *models.MarketSearchRequest) ([]entities.Market, error)
		CreateSubPrefecture(subPrefecture *entities.SubPrefecture) error
		CreateDistrict(district *entities.District) error
		GetAllDistrict(subprefecture uint64) ([]entities.District, error)
		GetAllSubPrefecture() ([]entities.SubPrefecture, error)
	}
)
