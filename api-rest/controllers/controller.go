package controllers

import (
	"net/http"
	"rest-api/connection"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func HomeDefault(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"home": "Welcome",
	})
}

func FindByFullName(c *gin.Context) {
	var name models.Name
	search := c.Param("name")

	connection.DATABASE.Where(&models.Name{Name: search}).First(&name)
	if name.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": search,
		})

		return

	}

	c.JSON(http.StatusOK, name)
}

func GetAllNames(c *gin.Context) {
	var names []models.Name
	connection.DATABASE.Find(&names)
	c.JSON(http.StatusOK, names)
}

func CreateNewInformation(c *gin.Context) {

	var name models.Name
	if err := c.ShouldBindJSON(&name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.CheckNamesStruct(&name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return

	}
	connection.DATABASE.Create(&name)
	c.JSON(http.StatusCreated, name)

}

func GetNameById(c *gin.Context) {
	var name models.Name
	connection.DATABASE.First(&name, c.Params.ByName("id"))
	if name.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"not found": "not found by id",
		})
		return

	}

	c.JSON(http.StatusOK, name)
}

func DeleteById(c *gin.Context) {
	var name models.Name
	connection.DATABASE.Delete(&name, c.Params.ByName("id"))
	c.JSON(http.StatusOK, gin.H{
		"message": "record deleted successfully",
	})
	c.Redirect(http.StatusMovedPermanently, "/api/names")
}

func EditById(c *gin.Context) {
	var name models.Name
	connection.DATABASE.First(&name, c.Params.ByName("id"))
	if err := c.ShouldBindJSON(&name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.CheckNamesStruct(&name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return

	}

	connection.DATABASE.Model(&name).UpdateColumns(name)

	c.JSON(http.StatusOK, gin.H{

		"data":    name,
		"message": "updated sucessfully",
	})
}
