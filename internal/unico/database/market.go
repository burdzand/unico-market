package database

import (
	"errors"
	"fmt"
	"unico/internal/unico/entities"
	"unico/internal/unico/models"
	logger "unico/pkg/log"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

var ErrAlreadyExist = errors.New("Record Already Exist")

var log *logger.Logger

func init() {
	log = logger.NewLogger()
}

type (
	DatabaseHandle struct {
		db *gorm.DB
	}
)

func NewDatabaseHandle(db *gorm.DB) *DatabaseHandle {
	return &DatabaseHandle{db: db}
}

func (r *DatabaseHandle) GetDB() *gorm.DB {
	return r.db
}

// GetMarket
func (r *DatabaseHandle) GetMarket(id uuid.UUID) (*entities.Market, error) {
	market := &entities.Market{}
	err := r.db.Model(&entities.Market{}).Preload("SubPrefecture").Preload("District").Where("id = ?", id).First(market).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return market, nil
}

// GetMarketByRecord
func (r *DatabaseHandle) GetMarketByRecord(record string) (*entities.Market, error) {
	market := &entities.Market{}
	err := r.db.Model(&entities.Market{}).Preload("SubPrefecture").Preload("District").Where("record = ?", record).First(market).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return market, nil
}

// GetAllMarket
func (r *DatabaseHandle) GetAllMarket(search *models.MarketSearchRequest) ([]entities.Market, error) {
	var markets = []entities.Market{}

	db := r.db.Model(&entities.Market{}).Joins("SubPrefecture").Joins("District")
	fmt.Println("PASSEi")
	if search.District != "" {
		db = db.Where("\"District\".\"name\" ILIKE ?", "%"+search.District+"%")
	}

	if search.Neighborhood != "" {
		db = db.Where("\"markets\".\"neighborhood\" ILIKE ?", "%"+search.Neighborhood+"%")
	}

	if search.Zone != "" {
		db = db.Where("\"SubPrefecture\".\"zone\" ILIKE ?", "%"+search.Zone+"%")
	}

	if search.Name != "" {
		db = db.Where("\"markets\".\"name\" ILIKE ?", "%"+search.Name+"%")
	}

	err := db.Find(&markets).Error
	return markets, err
}

// CreateMarket
func (r *DatabaseHandle) CreateMarket(market *entities.Market) error {
	return r.db.Create(market).Error
}

// Update
func (r *DatabaseHandle) UpdateMarket(market *entities.Market) error {
	return r.db.Model(&entities.Market{}).Where("id = ?", market.ID).Updates(market).Error
}

// RemoveMarket
func (r *DatabaseHandle) RemoveMarket(record string) error {
	return r.db.Where("record = ?", record).Delete(&entities.Market{}).Error
}
