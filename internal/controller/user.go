// controllers/user_controller.go
package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nsaltun/bitirici/internal/model"
	"github.com/nsaltun/bitirici/internal/service"
)

type userController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		userService: service,
	}
}

func (u *userController) CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	createdUser, err := u.userService.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(createdUser)
}

func (u *userController) GetUsers(c *fiber.Ctx) error {
	users, err := u.userService.GetUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(users)
}

func (u *userController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := u.userService.GetUser(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(user)
}

func (u *userController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateUser model.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	updatedUser, err := u.userService.UpdateUser(id, updateUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(updatedUser)
}

func (u *userController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := u.userService.DeleteUser(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
