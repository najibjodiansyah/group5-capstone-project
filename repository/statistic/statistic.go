package statistic

import (
	"capstone-project/entities"
	"database/sql"
)

type StatisticRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *StatisticRepository {
	return &StatisticRepository{db: db}
}

func (r *StatisticRepository) Get() (entities.Statistic, error) {
	var result *sql.Rows
	var err error
	var statistic entities.Statistic
	available := "tersedia"
	maintenance := "pemeliharaan"
	used := "digunakan"
	result, err = r.db.Query(`select count(id) from items`)
	if err != nil {
		return statistic, err
	}

	for result.Next() {
		err := result.Scan(&statistic.TotalAsset)
		if err != nil {
			return statistic, err
		}
	}
	result, err = r.db.Query(`select count(id) from items where availableStatus=?`, used)
	if err != nil {
		return statistic, err
	}

	for result.Next() {
		err := result.Scan(&statistic.TotalUsed)
		if err != nil {
			return statistic, err
		}
	}
	result, err = r.db.Query(`select count(id) from items where availableStatus=?`, available)
	if err != nil {
		return statistic, err
	}

	for result.Next() {
		err := result.Scan(&statistic.TotalAvailable)
		if err != nil {
			return statistic, err
		}
	}
	result, err = r.db.Query(`select count(id) from items where availableStatus=?`, maintenance)
	if err != nil {
		return statistic, err
	}

	for result.Next() {
		err := result.Scan(&statistic.TotalMaintenance)
		if err != nil {
			return statistic, err
		}
	}
	return statistic, nil
}
