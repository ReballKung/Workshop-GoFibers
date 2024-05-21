package main

import (
	"fmt"
	"go-fiber-test/database"
	m "go-fiber-test/models"      // เรียกใช้ Models
	router "go-fiber-test/routes" // เรียกใช้ Router

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// * Connection Database
func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",        // Username
		"",            // password
		"127.0.0.1",   // ip address
		"3306",        // port
		"golang_test", // database name
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	// * เรียก Models เพื่อสร้าง Tables
	database.DBConn.AutoMigrate(&m.Dogs{})
	database.DBConn.AutoMigrate(&m.Company{})
	database.DBConn.AutoMigrate(&m.Users{})
}

func main() {
	app := fiber.New()
	initDatabase() // ! Use Connect Database
	router.InetRoutes(app)
	app.Listen(":3000")
}
