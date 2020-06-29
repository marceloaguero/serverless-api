package users

import (
	"database/sql"
	"os"
)

// MysqlDBRepo implements usecase repository
type MysqlDBRepo struct {
	db        *sql.DB
	tableName string
}

// NewMysqlRepo creates the repo
func NewMysqlRepo(dbURI, dbName, tableName string) *MysqlDBRepo {
	db, err := mysqlConnect(dbURI, dbName)
	if err != nil {
		os.Exit(1)
	}

	return &MysqlDBRepo{
		db:        db,
		tableName: tableName,
	}
}

func mysqlConnect(dbURI, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbURI+"/"+dbName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
