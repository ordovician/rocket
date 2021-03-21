package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func main() {
	pwd, _ := os.Getwd()
	log.Println("Current working directory: ", pwd)

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
