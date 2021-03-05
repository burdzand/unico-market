package model

import (
	"time"

	"github.com/gofrs/uuid"
	gormv2 "gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

func (u *Base) BeforeCreate(tx *gormv2.DB) (err error) {
	if u.ID == uuid.Nil {
		idUUID, err := uuid.NewV4()
		if err != nil {
			return err
		}
		u.ID = idUUID
	}
	return
}
