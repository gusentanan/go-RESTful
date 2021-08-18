package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gusentanan/go-RESTful/pkg/routes"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("unable to connect envronment variables: %s", err.Error())
	}
}

func getAddress() string {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	return fmt.Sprintf("%s:%s", host, port)
}

func listenToHttpReq(address string, mux *mux.Router) {
	log.Printf("start to listen at : %s", address)
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatalf("unable to listen to the network : %s", err.Error())
	}
}

func main() {

	loadEnv()
	address := getAddress()
	mux := routes.Router()

	listenToHttpReq(address, mux)

}
