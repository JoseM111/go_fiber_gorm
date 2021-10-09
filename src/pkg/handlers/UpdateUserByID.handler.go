package handlers

import (
	"github.com/gofiber/fiber/v2"
	
	. "go_fiber_gorm/src/pkg/database"
	. "go_fiber_gorm/src/pkg/models"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func UpdateUserByID(ctx *fiber.Ctx) error {
	// ™ ____________
	id := ctx.Params("id")
	user := new(User)
	
	DB.First(&user, id)
	
	if user.Email == "" {
		return ctx.Status(500).SendString("User not found...")
	}
	
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	
	DB.Save(&user)
	return ctx.JSON(&user)
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰








