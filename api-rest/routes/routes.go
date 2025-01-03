package routes

import (
	"rest-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	router := gin.Default()
	router.Use(

		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE"},
			AllowHeaders: []string{"Content-Type"},
		}),
	)

	router.SetTrustedProxies(nil)

	router.GET("/api/v1", controllers.HomeDefault)
	router.GET("/api/v1/names", controllers.GetAllNames)
	router.GET("/api/v1/names/:id", controllers.GetNameById)
	router.GET("/api/v1/names/findby/:name", controllers.FindByFullName)
	router.POST("/api/v1/names", controllers.CreateNewInformation)
	router.DELETE("/api/v1/names/:id", controllers.DeleteById)
	router.PATCH("/api/v1/names/:id", controllers.EditById)

	router.Run(":8083")

}
