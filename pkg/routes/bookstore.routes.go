package routes

import (
	controller "github.com/eslamabdallah74/go_book_management/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStore = func(router *mux.Router) {
	router.HandleFunc("/books/", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}
