package commons

import (
	"log"

	"github.com/Smoked22/api-go-mysql/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@tcp(localhost:3306)/pruebas?charset=utf8")

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.DB()
	log.Println("Migrando...")

	db.AutoMigrate(&models.Persona{})
}
