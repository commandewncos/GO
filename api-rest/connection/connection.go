package connection

import (
	"rest-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DATABASE *gorm.DB
	ERROR    error
)

func ConnectionWithDatabase() {
	connectionString := "application:appPasSworD@tcp(172.20.0.2:3306)/names?charset=utf8mb4&parseTime=True&loc=Local"

	DATABASE, ERROR = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if ERROR != nil {
		panic(ERROR.Error())
	}

	DATABASE.AutoMigrate(&models.Name{})
}
