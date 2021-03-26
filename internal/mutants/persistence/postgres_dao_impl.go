package persistence

import (
	"database/sql"
	"errors"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/domain"
)

type dao struct {
	tableName string
	databaseConnectionProvider func() (*sql.DB, error)
}

func New(tableName string, databaseConnectionProvider func() (*sql.DB, error)) domain.Dao {
	return &dao{tableName: tableName, databaseConnectionProvider: databaseConnectionProvider}
}

func (dao *dao)AddAnalysis(dna string, isMutant bool) (bool, error) {
	query := `INSERT INTO ` + dao.tableName + `(dna, isMutant) VALUES ($1, $2)`
	db, err := dao.databaseConnectionProvider()
	if err != nil {
		return false, err
	}
	defer db.Close()

	statement, err := db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer statement.Close()

	rows, err := statement.Exec(dna, isMutant)
	if err != nil {
		return false, err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return false, err
	} else if rowsAffected != 1 {
		return false, errors.New("Expected one affected row")
	}
	return isMutant, nil
}
