package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vishalc412/go-bookstore-crud/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(router)
	http.Handle("/", router)
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
