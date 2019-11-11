package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func homeLink(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	s3Bucket := os.Getenv("S3_BUCKET")
	//secretKey := os.Getenv("SECRET_KEY")

	fmt.Fprintf(w, s3Bucket)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
