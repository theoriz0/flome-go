package database

import (
	"fmt"

	"github.com/theoriz0/flome-go/internal/config"
	"github.com/theoriz0/flome-go/internal/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func Init(dc *config.DatabaseConfig) {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.Username, dc.Password, dc.Address, dc.Port, dc.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Init DB failed", log.String("err", err.Error()))
	}
	GlobalDB = db
}
