package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"

	"exchange/global"
)

func InitDB() {
	dsn := AppConfig.Database.Dns
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connt faile db")
	}
	sqldb, err := db.DB()

	sqldb.SetMaxIdleConns(10)
	sqldb.SetConnMaxLifetime(time.Hour)
	global.DB = db
}
