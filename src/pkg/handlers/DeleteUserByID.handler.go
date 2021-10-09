package handlers

import (
	"github.com/gofiber/fiber/v2"
	
	. "go_fiber_gorm/src/pkg/database"
	. "go_fiber_gorm/src/pkg/models"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func DeleteUserByID(ctx *fiber.Ctx) error {
	// ™ ____________
	id := ctx.Params(":id")
	var user User
	
	DB.First(&user, id)
	
	if user.Email == "" {
		return ctx.Status(500).SendString("User not found...")
	}
	
	DB.Delete(&user)
	return ctx.SendString("User has been deleted")
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
