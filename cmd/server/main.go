package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	log.Println("Running on port: ", port)
}
