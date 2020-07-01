package mysql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/marceloaguero/serverless-api/users"
	// Blank import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const timeout = time.Second * 5

// MysqlDBRepo implements usecase repository
type mysqlRepo struct {
	db        *sql.DB
	tableName string
}

// NewMysqlRepo creates the repo
func NewMysqlRepo(dsName, dbName, tableName string) (users.Repository, error) {
	db, err := mysqlConnect(dsName, dbName)
	if err != nil {
		return nil, err
	}

	return &mysqlRepo{
		db:        db,
		tableName: tableName,
	}, nil
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
func (r *mysqlRepo) Create(ctx context.Context, user *users.User) error {
	insertQuery := "INSERT INTO " + r.tableName +
		"(id, email, name, age) VALUES (?, ?, ?, ?)"
	stmt, err := r.db.PrepareContext(ctx, insertQuery)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user.ID, user.Email, user.Name, user.Age)
	return err
}

// Get a user
func (r *mysqlRepo) Get(ctx context.Context, id string) (*users.User, error) {
	result := users.User{}
	selectQuery := "SELECT id, email, name, age FROM " +
		r.tableName +
		" WHERE id = ?"

	err := r.db.QueryRowContext(ctx, selectQuery, id).Scan(
		result.ID, result.Email, result.Name, result.Age,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAll users
func (r *mysqlRepo) GetAll(ctx context.Context) ([]*users.User, error) {
	result := make([]*users.User, 0)
	selectQuery := "SELECT id, email, name, age FROM " +
		r.tableName

	rows, err := r.db.QueryContext(ctx, selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &users.User{}
		if err := rows.Scan(
			&user.ID, &user.Email, &user.Name, &user.Age,
		); err != nil {
			return nil, err
		}

		result = append(result, user)

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return result, nil
}

// Update a user
func (r *mysqlRepo) Update(ctx context.Context, id string, user *users.UpdateUser) error {
	updateQuery := "UPDATE " +
		r.tableName +
		" SET email = ?, name = ?, age = ? WHERE id = ?"
	stmt, err := r.db.PrepareContext(ctx, updateQuery)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user.Email, user.Name, user.Age, id)
	return err
}

// Delete a user
func (r *mysqlRepo) Delete(ctx context.Context, id string) error {
	deleteQuery := "DELETE FROM " +
		r.tableName +
		" WHERE id = ?"
	stmt, err := r.db.PrepareContext(ctx, deleteQuery)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, id)
	return err
}
