package main

import (
	"log"
	"net/http"

	"github.com/ImArnav/go-bookstore/pkg/routes" //all routes are absolute
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //need the package for connection
)

func main() {
	r := mux.NewRouter()              //mux router
	routes.RegisterBookStoreRoutes(r) //routes info

	http.Handle("/", r)                                //handle all routes
	log.Fatal(http.ListenAndServe("locahost:9010", r)) //ListenAndServer Just to create a server

}
