package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nsaltun/bitirici/internal/controller"
)

type Router interface {
	Setup()
}

type router struct {
	app            *fiber.App
	userController controller.UserController
}

func New(app *fiber.App, userController controller.UserController) Router {
	return &router{
		app,
		userController,
	}
}

func (r *router) Setup() {
	api := r.app.Group("/api")
	api.Post("/users", r.userController.CreateUser)
	api.Get("/users", r.userController.GetUsers)
	api.Get("/users/:id", r.userController.GetUser)
	api.Put("/users/:id", r.userController.UpdateUser)
	api.Delete("/users/:id", r.userController.DeleteUser)
}
