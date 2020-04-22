package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var SqlDB *gorm.DB

func init() {
	// Open file
	db, err := gorm.Open("sqlite3", "database/cpu.db")
	// Error
	if err != nil {
		panic(err)
	}
	SqlDB = db
	SqlDB.LogMode(true)

	// Creating the table
	if !SqlDB.HasTable(&Temps{}) {
		SqlDB.CreateTable(&Temps{})
		SqlDB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Temps{})
	}
}
