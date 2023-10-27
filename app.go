package main

import (
	"CrudORM/models"
	routes "CrudORM/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "usergo:12345@tcp(localhost:3306)/examplego?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Todo{})
	r := routes.SetupRouter(db)
	r.Run(":8080")
}
