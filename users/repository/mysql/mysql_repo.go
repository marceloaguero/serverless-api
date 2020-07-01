package mysql

import (
	"context"
	"database/sql"
	"fmt"
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
	conn := fmt.Sprintf("%s/%s", dsName, dbName)

	db, err := sql.Open("mysql", conn)
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
	insertQuery := fmt.Sprintf("INSERT INTO %s (id, email, name, age) VALUES (?, ?, ?, ?)", r.tableName)

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

	selectQuery := fmt.Sprintf("SELECT id, email, name, age FROM %s WHERE id = ?", r.tableName)

	err := r.db.QueryRowContext(ctx, selectQuery, id).Scan(
		&result.ID, &result.Email, &result.Name, &result.Age,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAll users
func (r *mysqlRepo) GetAll(ctx context.Context) ([]*users.User, error) {
	result := make([]*users.User, 0)

	selectQuery := fmt.Sprintf("SELECT id, email, name, age FROM %s", r.tableName)

	rows, err := r.db.QueryContext(ctx, selectQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := &users.User{}

		err := rows.Scan(
			&user.ID, &user.Email, &user.Name, &user.Age,
		)
		if err != nil {
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
	updateQuery := fmt.Sprintf("UPDATE %s SET email = ?, name = ?, age = ? WHERE id = ?", r.tableName)

	stmt, err := r.db.PrepareContext(ctx, updateQuery)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user.Email, user.Name, user.Age, id)

	return err
}

// Delete a user
func (r *mysqlRepo) Delete(ctx context.Context, id string) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = ?", r.tableName)

	stmt, err := r.db.PrepareContext(ctx, deleteQuery)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, id)

	return err
}
