package main

import (
	"server/database"
	"server/routes"
)

func main() {
	connectionString := database.FormatConnectionString()
	database.Connect(connectionString)
	database.Migrate()
	routes.Run()
}
