package main

import (
	"myapp/config"
	"myapp/lib/database"
	"myapp/routes"
)

func main() {
	database.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(config.PORT))
}
