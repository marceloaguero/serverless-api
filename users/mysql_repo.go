package users

import (
	"context"
	"database/sql"
	"os"
	"time"

	// Blank import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const timeout = time.Second * 5

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

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
