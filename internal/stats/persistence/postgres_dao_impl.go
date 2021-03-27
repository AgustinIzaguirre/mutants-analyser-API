package persistence

import (
	"database/sql"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/stats/domain"
	"log"
)

type dao struct {
	tableName string
	databaseConnectionProvider func() (*sql.DB, error)
}

func New(tableName string, databaseConnectionProvider func() (*sql.DB, error)) domain.Dao {
	return &dao{tableName: tableName, databaseConnectionProvider: databaseConnectionProvider}
}

func (dao *dao) GetStats() (domain.Stats, errors.ApiError) {
	query := `SELECT sum(case when IsMutant then 1 else 0 end) AS Mutants,` +
			`sum(case when not IsMutant then 1 else 0 end) AS Human FROM ` + dao.tableName + `;`
	db, rows, err := dao.makeQuery(query)
	if err != nil {
		log.Fatal(err)
		return domain.Stats{}, errors.NewInternalServerError(err.Error())
	}
	defer db.Close()
	defer rows.Close()

	var result domain.Stats
	for rows.Next() {
		err := rows.Scan(&result.Mutants, &result.Humans)
		if err != nil {
			return domain.Stats{}, errors.NewInternalServerError(err.Error())
		}
	}
	if result.Mutants == 0 {
		result.Ratio = 0
	} else if result.Humans == 0 {
		result.Ratio = float64(result.Mutants)
	} else {
		result.Ratio = float64(result.Mutants) / float64(result.Humans)
	}
	return result, nil
}

func (dao *dao) makeQuery(query string) (*sql.DB, *sql.Rows, error) {
	db, err := dao.databaseConnectionProvider()
	if err != nil {
		return db, &sql.Rows{}, err
	}
	rows, queryErr := db.Query(query)
	if queryErr != nil {
		return db, rows, err
	}
	return db, rows, nil
}