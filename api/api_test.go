package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"unico/internal/unico/database"
	"unico/internal/unico/entities"
	"unico/internal/unico/migrations"
	"unico/internal/unico/models"
	"unico/internal/unico/services"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var server *gin.Engine
var db *database.DatabaseHandle

func init() {
	server = gin.Default()
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
	ConfigureRoutes(server, db)

}

func createDistrict() *entities.District {
	subPrefecture := &entities.SubPrefecture{}
	subPrefecture.Name = gofakeit.Name()
	subPrefecture.Zone = entities.East
	err := db.GetDB().Save(subPrefecture).Error
	if err != nil {
		return nil
	}

	district := &entities.District{}
	district.Name = gofakeit.Name()
	district.SubPrefectureID = subPrefecture.ID
	err = db.CreateDistrict(district)
	if err != nil {
		return nil
	}
	return district
}

func getMarketRequest() *models.MarketCreateRequest {
	district := createDistrict()
	if district == nil {
		return nil
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
	return request

}

func TestNpRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/norount", nil)
	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestHttpCreateMarket(t *testing.T) {
	w := httptest.NewRecorder()

	request := getMarketRequest()
	jsonValue, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", "/api/v1/market/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}
	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHttpUpdateMarket(t *testing.T) {
	w := httptest.NewRecorder()

	marketreq := getMarketRequest()
	market, err := services.CreateMarket(marketreq, db)
	if err != nil {
		t.Fatal(err)
	}

	request := &models.MarketUpdateRequest{}
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
	request.SubPrefectureID = market.SubPrefectureID
	request.DistrictID = market.DistrictID

	jsonValue, _ := json.Marshal(request)
	req, err := http.NewRequest("PUT", "/api/v1/market/"+market.Record, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}
	server.ServeHTTP(w, req)
	fmt.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHttpDeletearket(t *testing.T) {
	w := httptest.NewRecorder()

	marketreq := getMarketRequest()
	market, err := services.CreateMarket(marketreq, db)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "/api/v1/market/"+market.Record, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}
	server.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
