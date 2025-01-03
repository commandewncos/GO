package router

import (
	"net/http"

	"github.com/Wilson-cmd/GO/controller"
)

func GetRoutes() {

	http.HandleFunc("/", controller.GetAllNames)
	http.HandleFunc("/create", controller.InsertNames)
	http.HandleFunc("/insert", controller.GetInfoByForm)
	http.HandleFunc("/delete", controller.DeleteById)
	http.HandleFunc("/edit", controller.EditeById)
	http.HandleFunc("/update", controller.Update)

}
