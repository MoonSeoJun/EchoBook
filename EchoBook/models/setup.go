package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB define
var DB *gorm.DB

// ConnectDataBase connect mysql
func ConnectDataBase() {
	dsn := "root:password@tcp(host:port)/bookstore"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Book{})

	DB = db
}
