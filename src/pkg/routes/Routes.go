package routes

import (
	"github.com/gofiber/fiber/v2"
	
	. "go_fiber_gorm/src/pkg/handlers"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

// Routes router
func Routes(app *fiber.App) {
	// ™ ____________
	// GET request: GetUsers
	app.Get("/users", GetUsers)
	// GET request: GetUserByID
	app.Get("/user/:id", GetUserByID)
	// POST request: SaveUser
	app.Post("/user", SaveUser)
	// PUT request: UpdateUserByID
	app.Put("/user/:id", UpdateUserByID)
	// DELETE request: DeleteUserByID
	app.Delete("/user/:id", DeleteUserByID)
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
