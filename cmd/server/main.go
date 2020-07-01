package main

import (
	"log"
	"net/http"
	"os"

	"github.com/marceloaguero/serverless-api/users"
	mysqlrepo "github.com/marceloaguero/serverless-api/users/repository/mysql"
	transport "github.com/marceloaguero/serverless-api/users/transport/http"
)

func main() {
	port := os.Getenv("PORT")
	dsName := os.Getenv("DB_DSN")
	dbName := os.Getenv("DB_NAME")
	tableName := os.Getenv("TABLE_NAME")

	// repository := NewMongoRepo(dbURI, dbName, tableName)
	repository, err := mysqlrepo.NewMysqlRepo(dsName, dbName, tableName)
	// NewMysqlRepo(dsName, dbName, tableName)
	if err != nil {
		log.Panic(err)
	}

	usecase := users.NewUsecase(repository)
	router, err := transport.Routes(usecase)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Running on port: ", port)
	log.Panic(http.ListenAndServe(":"+port, router))
}
