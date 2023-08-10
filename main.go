package main

import (
	"NEWGOLANG/prodService"
	"log"
	"net/http"
)

func main() {
	prodService.Init()
	log.Println("Starting server on port 6068")
	log.Println(http.ListenAndServe(":6068", nil))
}
