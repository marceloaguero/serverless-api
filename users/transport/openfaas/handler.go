package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	fmt.Println("Como dice que le va?")
	port := os.Getenv("PORT")
	dsName := os.Getenv("DB_DSN")
	dbName := os.Getenv("DB_NAME")
	tableName := os.Getenv("TABLE_NAME")

	fmt.Println(port, dsName, dbName, tableName)
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
