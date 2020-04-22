package models

import (
	_ "errors"
	_ "fmt"
)

type Temps struct {
	Id          int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Tdatetime   string `gorm:"not null" form:"tdatetime" json:"tdatetime"`
	Temperature string `gorm:"not null" form:"temperature" json:"temperature"`
	Fan         string `gorm:"not full" form:"fan" json:"fan"`
}

func ListInfo() []Temps {
	var temps []Temps
	// SELECT * FROM temps
	//fmt.Printf("get temps...\n")
	SqlDB.Find(&temps)
	return temps
}
