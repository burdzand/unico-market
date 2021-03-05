package entities

import (
	"unico/pkg/database/model"
)

type Zone string

const (
	East  Zone = "Lest"
	West  Zone = "Oeste"
	North Zone = "Norte"
	South Zone = "Sul"
)

type SubPrefecture struct {
	ID   uint64 `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index"`
	Zone Zone   `json:"zone"`
}

type District struct {
	ID              uint64        `json:"id,omitempty" gorm:"primaryKey"`
	Name            string        `json:"name" gorm:"index"`
	SubPrefectureID uint64        `json:"subPrefectureID,omitempty"`
	SubPrefecture   SubPrefecture `json:"-"`
}

type Market struct {
	model.Base
	Name            string        `json:"name" gorm:"index"`
	Record          string        `json:"record" gorm:"unique"`
	Latitude        string        `json:"latitude"`
	Longitude       string        `json:"longitude"`
	SetCents        uint64        `json:"setcents"`
	Area            uint64        `json:"area"`
	SubPrefectureID uint64        `json:"subPrefectureID,omitempty"`
	SubPrefecture   SubPrefecture `json:"subPrefecture,omitempty"`
	DistrictID      uint64        `json:"districtID,omitempty"`
	District        District      `json:"district,omitempty"`
	Zone8           string        `json:"zone8,omitempty"`
	Street          string        `json:"Street,omitempty"`
	Number          string        `json:"Number,omitempty"`
	Neighborhood    string        `json:"neighborhood,omitempty"`
	Reference       string        `json:"reference,omitempty"`
}
