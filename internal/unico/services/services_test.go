package services

import (
	"errors"
	"os"
	"testing"
	"unico/internal/unico/database"
	"unico/internal/unico/entities"
	"unico/internal/unico/migrations"
	"unico/internal/unico/models"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *database.DatabaseHandle

func init() {
	gofakeit.Seed(0)
	connectiondb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DB_CONNECTION"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to database: ", err)
	}

	err = migrations.Run(connectiondb)
	if err != nil {
		log.Panic("Failed to run migration: ", err)
	}

	db = database.NewDatabaseHandle(connectiondb)
}
func createSubPrefecture() *entities.SubPrefecture {
	subPrefecture := &entities.SubPrefecture{}
	subPrefecture.Name = gofakeit.Name()
	subPrefecture.Zone = entities.East
	err := db.GetDB().Save(subPrefecture).Error
	if err != nil {
		return nil
	}
	return subPrefecture
}

func createDistrict() *entities.District {
	subPrefecture := createSubPrefecture()
	district := getDistrict()
	district.SubPrefectureID = subPrefecture.ID
	err := db.CreateDistrict(district)
	if err != nil {
		return nil
	}
	return district
}

func getSubPrefecture() *entities.SubPrefecture {
	subPrefecture := &entities.SubPrefecture{}
	subPrefecture.Name = gofakeit.Name()
	subPrefecture.Zone = entities.East
	return subPrefecture
}

func getDistrict() *entities.District {
	district := &entities.District{}
	district.Name = gofakeit.Name()
	return district
}

func geMarket() *entities.Market {
	request := &entities.Market{}
	request.Name = gofakeit.Name()
	request.Latitude = gofakeit.DigitN(10)
	request.Longitude = gofakeit.DigitN(10)
	request.SetCents = gofakeit.Uint64()
	request.Area = gofakeit.Uint64()
	request.Zone8 = gofakeit.BeerName()
	request.Street = gofakeit.Name()
	request.Number = "s/n"
	request.Neighborhood = gofakeit.Name()
	request.Reference = gofakeit.Name()
	request.Record = gofakeit.DigitN(10)
	return request
}

func createMarket() *entities.Market {
	district := createDistrict()
	if district == nil {
		return nil
	}
	market := geMarket()
	market.SubPrefectureID = district.SubPrefectureID
	market.DistrictID = district.ID

	err := db.CreateMarket(market)
	if err != nil {
		return nil
	}
	return market
}

func TestCreateSubPrefecture(t *testing.T) {
	subPrefecture := getSubPrefecture()
	err := db.CreateSubPrefecture(subPrefecture)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateDistrict(t *testing.T) {

	subPrefecture := createSubPrefecture()
	if subPrefecture == nil {
		t.Fatal(errors.New("No SubPrefecture"))
	}

	district := getDistrict()
	district.SubPrefectureID = subPrefecture.ID
	err := db.CreateDistrict(district)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateMarket(t *testing.T) {

	district := createDistrict()
	if district == nil {
		t.Fatal(errors.New("No District"))
	}

	request := &models.MarketCreateRequest{}

	request.Name = gofakeit.Name()
	request.Latitude = gofakeit.DigitN(10)
	request.Longitude = gofakeit.DigitN(10)
	request.SetCents = gofakeit.Uint64()
	request.Area = gofakeit.Uint64()
	request.SubPrefectureID = district.SubPrefectureID
	request.DistrictID = district.ID
	request.Zone8 = gofakeit.BeerName()
	request.Street = gofakeit.Name()
	request.Number = "s/n"
	request.Neighborhood = gofakeit.Name()
	request.Reference = gofakeit.Name()
	request.Record = gofakeit.DigitN(10)

	_, err := CreateMarket(request, db)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveMarket(t *testing.T) {
	market := createMarket()
	err := db.RemoveMarket(market.Record)
	if err != nil {
		t.Fatal(err)
	}
	oldmarket, err := db.GetMarket(market.ID)
	if oldmarket != nil || err != nil {
		t.Fatal("Erro delete Market")
	}
}
