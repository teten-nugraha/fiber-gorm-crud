package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/teten-nugraha/golang-crud-fiber/database"
	"github.com/teten-nugraha/golang-crud-fiber/routes"
)

func main() {
	// create new server
	server := fiber.New()

	// connecting db
	dbErr := database.InitDatabase()
	if dbErr != nil {
		panic(dbErr)
	}

	// allowed CORS
	server.Use(cors.New())

	// add logger
	server.Use(logger.New())

	// setup routes
	routes.Setup(server)

	// listen in port 4000
	server.Listen(":4000")
}
