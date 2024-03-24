package main

import (
	"final-project/pkg/database"
	"final-project/pkg/routers"
)

func main() {
	port := ":5050"

	gorm:= database.DataBaseInit()

	routers.StartServer(gorm).Run(port)

	
}

