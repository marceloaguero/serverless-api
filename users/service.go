package users

import (
	"os"
)

// Init sets up an instance of this domains
// usecase, pre-configured with the dependencies.
func Init() (Usecase, error) {
	dsName := os.Getenv("DB_DSN")
	dbName := os.Getenv("DB_NAME")
	tableName := os.Getenv("TABLE_NAME")

	// repository := NewMongoRepo(dbURI, dbName, tableName)
	repository, err := NewMysqlRepo(dsName, dbName, tableName)
	if err != nil {
		return nil, err
	}

	usecase := NewUsecase(repository)

	return usecase, nil
}
