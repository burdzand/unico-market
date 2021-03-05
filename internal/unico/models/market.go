package models

import (
	"unico/internal/unico/entities"
)

type MarketSearchRequest struct {
	Name         string `form:"name"`
	Zone         string `form:"zone"`
	District     string `form:"district"`
	Neighborhood string `form:"neighborhood"`
}

type MarketUpdateRequest struct {
	Name            string `json:"name" validate:"required"`
	Latitude        string `json:"latitude" validate:"required"`
	Longitude       string `json:"longitude" validate:"required"`
	SetCents        uint64 `json:"setcents" validate:"required"`
	SubPrefectureID uint64 `json:"subPrefectureID,omitempty" validate:"required"`
	DistrictID      uint64 `json:"districtID,omitempty" validate:"required"`
	Zone8           string `json:"zone8,omitempty" validate:"required"`
	Street          string `json:"street" validate:"required"`
	Number          string `json:"number"`
	Neighborhood    string `json:"neighborhood" validate:"required"`
	Reference       string `json:"reference" validate:"required"`
	Area            uint64 `json:"area" validate:"required"`
}

func (mc *MarketUpdateRequest) ParseModel(ec *entities.Market) {
	ec.Name = mc.Name
	ec.Latitude = mc.Latitude
	ec.Longitude = mc.Longitude
	ec.SetCents = mc.SetCents
	ec.SubPrefectureID = mc.SubPrefectureID
	ec.DistrictID = mc.DistrictID
	ec.Zone8 = mc.Zone8
	ec.Street = mc.Street
	ec.Number = mc.Number
	ec.Neighborhood = mc.Neighborhood
	ec.Reference = mc.Reference
}

type MarketCreateRequest struct {
	Record string `json:"record" validate:"required"`
	MarketUpdateRequest
}

func (mc *MarketCreateRequest) ParseModel(ec *entities.Market) {
	ec.Name = mc.Name
	ec.Record = mc.Record
	ec.Latitude = mc.Latitude
	ec.Longitude = mc.Longitude
	ec.SetCents = mc.SetCents
	ec.SubPrefectureID = mc.SubPrefectureID
	ec.DistrictID = mc.DistrictID
	ec.Zone8 = mc.Zone8
	ec.Street = mc.Street
	ec.Number = mc.Number
	ec.Neighborhood = mc.Neighborhood
	ec.Reference = mc.Reference
}
