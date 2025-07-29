package main

import (
	"github.com/farinchan/thesis-attendance-backend/config"
	"github.com/farinchan/thesis-attendance-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	// config.DB.AutoMigrate(&model.User{})

	app.Get("/server", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Server is running"})
	})
	app.Post("/checkin", handler.AttendanceCheckin)

	app.Listen(":3000")
}
