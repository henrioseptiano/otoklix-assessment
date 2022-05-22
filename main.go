package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"otoklix-assessment.com/routes"
)

func main() {
	//env := os.Getenv("ENV")
	port := os.Getenv("PORT")

	app := fiber.New()
	routes.PostRoutes(app)
	app.Listen(":" + port)
}
