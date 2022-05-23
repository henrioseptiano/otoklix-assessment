package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"otoklix-assessment.com/routes"
)

func main() {
	//env := os.Getenv("ENV")
	port := os.Getenv("PORT")
	DB := os.Getenv("DB")

	db, err := gorm.Open(sqlite.Open(DB), &gorm.Config{})
	if err != nil {
		log.Println("Cannot connect to SQLite db")
		return
	}
	app := fiber.New()
	routes.PostRoutes(app, db)
	app.Listen(":" + port)
}
