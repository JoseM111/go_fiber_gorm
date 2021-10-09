package handlers

import (
	"github.com/gofiber/fiber/v2"
	
	. "go_fiber_gorm/src/pkg/database"
	. "go_fiber_gorm/src/pkg/models"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

// GetUsers /* only returns an error if their is one */
func GetUsers(ctx *fiber.Ctx) error {
	// ™ ____________
	var users []User
	
	DB.Find(&users)
	return ctx.JSON(&users)
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰













