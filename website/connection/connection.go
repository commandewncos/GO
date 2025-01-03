package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	E  error
)

func ConnectionWithDatabase() {
	dsn := "user:pasSword@tcp(172.19.0.2:3306)/people?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "application:appPasSworD@tcp(172.20.0.2:3306)/names?charset=utf8mb4&parseTime=True&loc=Local"
	DB, E = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if E != nil {
		panic(E.Error())
	}

}
