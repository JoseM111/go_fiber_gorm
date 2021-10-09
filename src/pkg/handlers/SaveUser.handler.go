package handlers

import (
	"github.com/gofiber/fiber/v2"
	
	. "go_fiber_gorm/src/pkg/database"
	. "go_fiber_gorm/src/pkg/models"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func SaveUser(ctx *fiber.Ctx) error {
	// ™ ____________
	user := new(User)
	
	/*| BodyParser binds the request body to a struct. It supports
	    decoding the following content types based on the Content-Type
	    header & convert the body to the user struct data: |*/
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	
	/*| Save update value in database, if the value
	    doesn't have primary key, will insert it |*/
	DB.Create(&user)
	/*| JSON converts any interface or string to JSON. |*/
	return ctx.JSON(&user)
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰












