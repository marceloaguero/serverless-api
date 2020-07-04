package main

import (
	"log"
	"net/http"
	"os"

	"github.com/marceloaguero/serverless-api/users"
	transport "github.com/marceloaguero/serverless-api/users/delivery"
	mysqlrepo "github.com/marceloaguero/serverless-api/users/repository/mysql"
)

func main() {
	port := os.Getenv("PORT")
	dsName := os.Getenv("DB_DSN")
	dbName := os.Getenv("DB_NAME")
	tableName := os.Getenv("TABLE_NAME")
	pathPrefix := "/function/serverless-api/"

	repository, err := mysqlrepo.NewMysqlRepo(dsName, dbName, tableName)
	if err != nil {
		log.Panic(err)
	}

	usecase := users.NewUsecase(repository)
	router, err := transport.Routes(usecase, pathPrefix)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Running on port: ", port)
	log.Panic(http.ListenAndServe(":"+port, router))
}
