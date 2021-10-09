package utils

import (
	"log"
	
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

// ListeningPort listening at port:3000
func ListeningPort(app *fiber.App) {
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}

// EnvVariableLoader .env file loader
func EnvVariableLoader() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
