package handler

import (
    "github.com/gofiber/fiber/v2"
    "github.com/farinchan/thesis-attendance-backend/model"
    "github.com/farinchan/thesis-attendance-backend/service"
    "strconv"

    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
    user := new(model.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    if err := validate.Struct(user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    if err := service.CreateUser(user); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
    users, err := service.GetUsers()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    user, err := service.GetUser(uint(id))
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "User not found"})
    }
    return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    user, err := service.GetUser(uint(id))
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "User not found"})
    }

    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    if err := validate.Struct(user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    if err := service.UpdateUser(&user); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    if err := service.DeleteUser(uint(id)); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.SendStatus(204)
}
