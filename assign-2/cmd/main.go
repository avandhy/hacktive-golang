package main

import (
	"assign-2/pkg/database"
	"assign-2/pkg/routers"
)

func main() {
	port := ":5000"

	gorm := database.GormInit()

	routers.StartServer(gorm).Run(port)

}
