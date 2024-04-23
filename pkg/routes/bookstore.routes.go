package routes

import (
	controller "github.com/eslamabdallah74/go_book_management/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStore = func(router *mux.Router) {
	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE")
}
