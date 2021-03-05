package database

import (
	"unico/internal/unico/entities"

	"github.com/lib/pq"
)

// CreateSubPrefecture
func (r *DatabaseHandle) CreateSubPrefecture(subPrefecture *entities.SubPrefecture) error {
	err := r.db.Save(subPrefecture).Error
	if perr, ok := err.(*pq.Error); ok && perr.Code == "23505" {
		return ErrAlreadyExist
	}
	return err
}

// CreateDistrict
func (r *DatabaseHandle) CreateDistrict(district *entities.District) error {
	err := r.db.Save(district).Error
	if perr, ok := err.(*pq.Error); ok && perr.Code == "23505" {
		return ErrAlreadyExist
	}
	return err
}

// GetAllSubPrefecture
func (r *DatabaseHandle) GetAllSubPrefecture() ([]entities.SubPrefecture, error) {
	var records = []entities.SubPrefecture{}
	err := r.db.Model(&entities.SubPrefecture{}).Find(&records).Error
	return records, err
}

// GetAllDistrict
func (r *DatabaseHandle) GetAllDistrict(subPrefecture uint64) ([]entities.District, error) {
	var records = []entities.District{}
	err := r.db.Model(&entities.District{}).Where("sub_prefecture_id = ?", subPrefecture).Find(&records).Error
	return records, err
}
