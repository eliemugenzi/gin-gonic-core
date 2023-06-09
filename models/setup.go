package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
   database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

   if err != nil {
	panic("Failed to connect to the DB")
   }

   err = database.AutoMigrate(&Todo{})

   if err != nil {
	return
   }

   DB = database
}