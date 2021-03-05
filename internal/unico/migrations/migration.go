package migrations

import (
	"unico/internal/unico/entities"

	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {

	db.Debug()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		MigrationSubPrefecture(),
		MigrationDistrict(),
		MigrationMarket(),
	})

	err := m.Migrate()
	if err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
	return nil
}

func MigrationMarket() *gormigrate.Migration {
	m := &gormigrate.Migration{
		ID: "MigrationMarketTable",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&entities.Market{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("market")
		},
	}

	return m
}

func MigrationSubPrefecture() *gormigrate.Migration {
	m := &gormigrate.Migration{
		ID: "MigrationSubPrefectureTable",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&entities.SubPrefecture{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("sub_prefectures")
		},
	}

	return m
}

func MigrationDistrict() *gormigrate.Migration {
	m := &gormigrate.Migration{
		ID: "MigrationDistrictTable",
		Migrate: func(db *gorm.DB) error {
			return db.AutoMigrate(&entities.District{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("districts")
		},
	}

	return m
}
