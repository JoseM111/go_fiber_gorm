package database

import (
	. "fmt"
	"os"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
	. "go_fiber_gorm/src/pkg/models"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

var DB *gorm.DB
var err error
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func InitialMigration() {
	// ™ ____________
	// loading .env variable
	var DB_SOURCE = os.Getenv("DB_SOURCE")
	
	DB, err = gorm.Open(mysql.Open(DB_SOURCE), &gorm.Config{})

	if err != nil {
		Println(err.Error())
		panic("Cannot connect to database")
	}

	err := DB.AutoMigrate(&User{})
	if err != nil { return }
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
