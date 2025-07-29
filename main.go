package main

import (
    "golang_backend/config"
    "golang_backend/handler"
    "golang_backend/model"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    config.ConnectDB()
    // config.DB.AutoMigrate(&model.User{})

    app.Post("/users", handler.CreateUser)
    app.Get("/users", handler.GetUsers)
    app.Get("/users/:id", handler.GetUser)
    app.Put("/users/:id", handler.UpdateUser)
    app.Delete("/users/:id", handler.DeleteUser)

    app.Listen(":3000")
}
