package users

import (
	"context"
	"os"
)

// UserService is the top level signature of this service
type UserService interface {
	Create(ctx context.Context, user *User) error
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}

// Init sets up an instance of this domains
// usecase, pre-configured with the dependencies.
func Init() (UserService, error) {
	dbURI := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	tableName := os.Getenv("TABLE_NAME")

	repository := NewDBRepo(dbURI, dbName, tableName)

	usecase := &Usecase{Repository: repository}

	return usecase, nil
}
