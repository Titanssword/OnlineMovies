package main

import (
	"log"
	"net/http"
	"./pack"
)


func main() {
	router := pack.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}