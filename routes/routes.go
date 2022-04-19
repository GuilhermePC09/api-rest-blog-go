package routes

import (
	"log"
	"net/http"

	"github.com/GuilhermePC09/api-rest-blog-go/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts", controllers.ListPosts).Methods("GET")
	router.HandleFunc("/posts", controllers.EditPost).Methods("PUT")
	router.HandleFunc("/posts", controllers.DeletePost).Methods("DELETE")

	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.ListUsers).Methods("GET")
	router.HandleFunc("/users", controllers.EditUser).Methods("PUT")
	router.HandleFunc("/users", controllers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
