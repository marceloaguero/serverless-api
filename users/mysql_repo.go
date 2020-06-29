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
func NewMysqlRepo(dsName, dbName, tableName string) *MysqlDBRepo {
	db, err := mysqlConnect(dsName, dbName)
	if err != nil {
		os.Exit(1)
	}

	return &MysqlDBRepo{
		db:        db,
		tableName: tableName,
	}
}

func mysqlConnect(dsName, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsName+"/"+dbName)
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

// Create a user
func (r *MysqlDBRepo) Create(ctx context.Context, user *User) error {
	insertQuery := "INSERT INTO " + r.tableName + "(id, email, name, age) VALUES (?, ?, ?, ?)"
	stmt, err := r.db.PrepareContext(ctx, insertQuery)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user.ID, user.Email, user.Name, user.Age)
	return err
}

// Get a user
func (r *MysqlDBRepo) Get(ctx context.Context, id string) (*User, error) {
	return nil, nil
}

// GetAll users
func (r *MysqlDBRepo) GetAll(ctx context.Context) ([]*User, error) {
	return nil, nil
}

// Update a user
func (r *MysqlDBRepo) Update(ctx context.Context, id string, user *UpdateUser) error {
	return nil
}

// Delete a user
func (r *MysqlDBRepo) Delete(ctx context.Context, id string) error {
	return nil
}
