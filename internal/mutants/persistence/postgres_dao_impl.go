package persistence

import (
	"database/sql"
	"fmt"
	"github.com/AgustinIzaguirre/mutants-analyser-api/internal/mutants/domain"
	"log"
)

type dao struct {
	tableName string
	databaseConnectionProvider func() (*sql.DB, error)
}

func New(tableName string, databaseConnectionProvider func() (*sql.DB, error)) domain.Dao {
	return &dao{tableName: tableName, databaseConnectionProvider: databaseConnectionProvider}
}

func (dao *dao)AddAnalysis(isMutant bool) error {
	query := `SELECT COUNT(*) FROM ` + dao.tableName + ` WHERE IsMutant;`
	db, rows, err := dao.makeQuery(query)
	defer db.Close()
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return err
		}
	}
	fmt.Println(count)
	return nil
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