package persistence

import (
	"database/sql"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/errors"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/domain"
)

type dao struct {
	tableName string
	databaseConnectionProvider func() (*sql.DB, error)
}

func New(tableName string, databaseConnectionProvider func() (*sql.DB, error)) domain.Dao {
	return &dao{tableName: tableName, databaseConnectionProvider: databaseConnectionProvider}
}

func (dao *dao)AddAnalysis(dna string, isMutant bool) (bool, errors.ApiError) {
	query := `INSERT INTO ` + dao.tableName + `(dna, isMutant) VALUES ($1, $2)`
	db, statement, err := dao.prepareQuery(query)
	if err != nil {
		return false, errors.NewInternalServerError(err.Error())
	}
	defer db.Close()
	defer statement.Close()

	rows, err := statement.Exec(dna, isMutant)
	if err != nil {
		return false, errors.NewInternalServerError(err.Error())
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return false, errors.NewInternalServerError(err.Error())
	} else if rowsAffected != 1 {
		return false, errors.NewInternalServerError("Expected one affected row")
	}
	return isMutant, nil
}

func (dao *dao) HasDNASequence(dna string) (bool, errors.ApiError) {
	query := `SELECT dna FROM ` + dao.tableName + ` where dna LIKE $1`
	db, statement, err := dao.prepareQuery(query)
	if err != nil {
		return false, errors.NewInternalServerError(err.Error())
	}
	defer db.Close()
	defer statement.Close()

	rows, err := statement.Query(dna)
	if err != nil {
		return false, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var currentDna string
		err := rows.Scan(&currentDna)
		if err != nil {
			return false, errors.NewInternalServerError(err.Error())
		}
		count ++
	}
	return count > 0, nil
}

func (dao *dao) prepareQuery(query string) (*sql.DB, *sql.Stmt, error) {
	db, err := dao.databaseConnectionProvider()
	if err != nil {
		return &sql.DB{}, &sql.Stmt{}, err
	}
	statement, err := db.Prepare(query)
	if err != nil {
		return &sql.DB{}, &sql.Stmt{}, err
	}
	return db, statement, nil
}