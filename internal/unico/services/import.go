package services

import (
	"context"
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"strconv"
	"unico/internal/unico/database"
	"unico/internal/unico/entities"
)

type Result struct {
	last_value uint64
}

func ImportDataCSV(ctx context.Context, db database.Database, file *multipart.FileHeader) error {
	filetemp, err := file.Open()

	reader := csv.NewReader(filetemp)

	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()

	if err != nil {
		log.Errorf("Error to read file: %s", err)
		return err
	}

	var lastDistrictID uint64 = 0
	var lastSubPrefectureID uint64 = 0

	for _, line := range record[1:] {

		// Save subPrefecture
		id, _ := strconv.ParseUint(line[7], 10, 32)

		subPrefecture := &entities.SubPrefecture{}
		subPrefecture.ID = id
		subPrefecture.Name = line[8]
		subPrefecture.Zone = entities.Zone(line[9])

		err := db.CreateSubPrefecture(subPrefecture)
		if err != nil && err != database.ErrAlreadyExist {
			return err
		}

		log.WithField("Name", subPrefecture.Name).
			WithField("ID", subPrefecture.ID).
			Info("SubPrefecture Created")

		// Fixing Sequence from subPrefecture and District
		var result Result
		db.GetDB().Raw("SELECT last_value FROM sub_prefectures_id_seq").Scan(&result)
		// UPDATE
		if result.last_value <= subPrefecture.ID && lastSubPrefectureID < subPrefecture.ID {
			db.GetDB().Exec("ALTER SEQUENCE sub_prefectures_id_seq RESTART WITH ?;", subPrefecture.ID+1)
			lastSubPrefectureID = subPrefecture.ID
			log.WithField("SEQUENCE", subPrefecture.ID+1).
				WithField("OLD", result.last_value).
				WithField("ID", subPrefecture.ID).
				Info(" Updated sub_prefectures_id_seq SEQUENCE")
		}

		// Save District
		id, _ = strconv.ParseUint(line[5], 10, 32)
		district := &entities.District{}
		district.ID = id
		district.Name = line[6]
		district.SubPrefectureID = subPrefecture.ID

		err = db.CreateDistrict(district)
		if err != nil && err != database.ErrAlreadyExist {
			return err
		}

		log.WithField("Name", district.Name).
			WithField("ID", district.ID).
			Info("District Created")

		db.GetDB().Raw("SELECT last_value FROM districts_id_seq").Scan(&result)
		// UPDATE
		if result.last_value <= district.ID && lastDistrictID < district.ID {
			db.GetDB().Exec("ALTER SEQUENCE districts_id_seq RESTART WITH ?;", district.ID+1)
			lastDistrictID = district.ID
			log.WithField("SEQUENCE", district.ID+1).
				WithField("OLD", result.last_value).
				WithField("ID", district.ID).
				Info(" Updated districts_id_seq SEQUENCE")

		}

		fmt.Println("AQU 333")

		// Save Market
		setCents, _ := strconv.ParseUint(line[3], 10, 32)
		area, _ := strconv.ParseUint(line[4], 10, 32)

		market := &entities.Market{}
		market.Name = line[11]
		market.Latitude = line[1]
		market.Longitude = line[2]
		market.SetCents = setCents
		market.Area = area

		market.Zone8 = line[10]

		market.Record = line[12]
		market.Street = line[13]
		market.Number = line[14]
		market.Neighborhood = line[15]

		if len(line) == 17 {
			market.Reference = line[16]
		}

		market.DistrictID = district.ID
		market.SubPrefectureID = subPrefecture.ID

		err = db.CreateMarket(market)
		if err != nil {
			log.WithField("Name", market.Name).
				WithField("ID", market.ID).
				WithField("Message", err).
				Error("Erro trye Market Create")
			return err
		}
		log.WithField("Name", market.Name).
			WithField("ID", market.ID).
			Info("Market Created")

	}
	return nil
}
