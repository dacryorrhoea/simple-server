package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	routes(router)

	err = http.ListenAndServe("localhost:4444", router)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(router *httprouter.Router) {
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	router.GET("/", GetList)
}
