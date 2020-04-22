package main

import (
	. "pi-db/models"
	. "pi-db/routers"
)

func main() {
	defer SqlDB.Close()

	router := InitRouter()

	router.Run(":8080")
}
