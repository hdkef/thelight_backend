package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {

}

func main() {

	mux := http.NewServeMux()

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}

}
