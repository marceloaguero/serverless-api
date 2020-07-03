package function

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/marceloaguero/serverless-api/users"
	mysqlrepo "github.com/marceloaguero/serverless-api/users/repository/mysql"
)

func init() {
	dsName := os.Getenv("DB_DSN")
	dbName := os.Getenv("DB_NAME")
	tableName := os.Getenv("TABLE_NAME")

	repository, err := mysqlrepo.NewMysqlRepo(dsName, dbName, tableName)
	if err != nil {
		log.Panic(err)
	}

	usecase := users.NewUsecase(repository)
	fmt.Println(usecase)
}

// Handle is the input point in openfaas
func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		input = body
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello world, input was: %s", string(input))))
}
