package routes

import (
	"go/api/controllers"
	"go/api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleAllRequests() {

	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/names", controllers.GetNames).Methods("Get")
	r.HandleFunc("/api/names/{id}", controllers.GetById).Methods("Get")
	r.HandleFunc("/api/names", controllers.CreateName).Methods("POST")
	r.HandleFunc("/api/names/{id}", controllers.DeleteName).Methods("Delete")
	r.HandleFunc("/api/names/{id}", controllers.EditName).Methods("Put")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
