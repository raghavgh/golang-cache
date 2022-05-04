package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raghavgh/bookmanagement/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}
