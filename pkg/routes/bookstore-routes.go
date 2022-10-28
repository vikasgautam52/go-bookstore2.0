package routes

import(
	"github.com/gorilla/mux"
	"github.com/vikasgautam52/go-bookstore2.0/controllers"
)

var RegisterBookStoreroutes = func(router *mux.Router){
	router.HandleFunc("/books/", controllers.CreateBook).Methods("PUSH")
	router.HandleFunc("/books/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}