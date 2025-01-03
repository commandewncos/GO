package connection

import (
	"github.com/auth/command/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DATABASE *gorm.DB
	ERROR    error
)

func ConnectionWithDatabase() {
	connectionString := "application:appPasSworD@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"

	DATABASE, ERROR = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if ERROR != nil {
		panic(ERROR.Error())
	}

	DATABASE.AutoMigrate(&model.Users{})
}
