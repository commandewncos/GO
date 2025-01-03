package routers

import (
	"net/http"
	"token/controllers"
)

func HandlersRequests() {

	http.HandleFunc("/api/public", controllers.PublicHandler)
	http.HandleFunc("/api/auth", controllers.LoginHandler)
	http.HandleFunc("/api/secure", controllers.AuthMiddlewarew(controllers.SecureHandler))

	http.ListenAndServe(":8085", nil)
}
