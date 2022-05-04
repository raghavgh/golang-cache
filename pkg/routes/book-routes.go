package routes

import (
	"github.com/gorilla/mux"
	"github.com/raghavgh/bookmanagement/pkg/controllers"
)

const (
	delete = "DELETE"
	put    = "PUT"
	get    = "GET"
	post   = "POST"
)

// RegisterBookRoutes register all api routes
var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods(post)
	router.HandleFunc("/book/", controllers.GetBooks).Methods(get)
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods(get)
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods(put)
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods(delete)
}
