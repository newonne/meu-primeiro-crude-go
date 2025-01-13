package main

import (
	"crude-go/database"
	"crude-go/routes"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	r := routes.SetupRoutes()
	r.Run(":8080") // Servidor rodando na porta 8080
}
