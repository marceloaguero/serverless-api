package users

import (
	"fmt"
	"os"
)

// Init sets up an instance of this domains
// usecase, pre-configured with the dependencies.
func Init() (Usecase, error) {
	dsName := os.Getenv("DB_DSN")
	// dbName := os.Getenv("DB_NAME")
	// tableName := os.Getenv("TABLE_NAME")

	fmt.Println(dsName)

	// repository := NewMongoRepo(dbURI, dbName, tableName)
	// repository, err := mysqlrepo.NewMysqlRepo(dsName, dbName, tableName)
	// NewMysqlRepo(dsName, dbName, tableName)
	// if err != nil {
	// 	return nil, err
	// }

	// usecase := NewUsecase(repository)

	// return usecase, nil
	return nil, nil
}
