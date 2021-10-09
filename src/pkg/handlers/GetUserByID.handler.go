package handlers

import (
	"github.com/gofiber/fiber/v2"
	
	. "go_fiber_gorm/src/pkg/database"
	. "go_fiber_gorm/src/pkg/models"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func GetUserByID(ctx *fiber.Ctx) error {
	// ™ ____________
	id := ctx.Params("id")
	var user User
	
	DB.Find(&user, id)
	return ctx.JSON(&user)
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰


















