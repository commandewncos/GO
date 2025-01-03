package router

import (
	"github.com/auth/command/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	router := gin.Default()
	router.Use(

		cors.New(cors.Config{
			// AllowCredentials: true,
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE"},
			AllowHeaders: []string{"Content-Type", "X-CSRF-Token"},
		}),
	)

	router.POST("/api/v1/protected", controllers.HandleProtected)
	router.POST("/api/v1/register", controllers.HandleRegister)
	router.POST("/api/v1/logout", controllers.HandleLogout)
	router.POST("/api/v1/login", controllers.HandleLogin)

	router.SetTrustedProxies(nil)
	router.Run(":3000")

}
