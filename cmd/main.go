// main.go
package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nsaltun/bitirici/internal/controller"
	"github.com/nsaltun/bitirici/internal/repository"
	"github.com/nsaltun/bitirici/internal/router"
	"github.com/nsaltun/bitirici/internal/service"
	"github.com/nsaltun/bitirici/lib/db"
	"github.com/nsaltun/bitirici/lib/logging"
)

func main() {
	logger := logging.New()
	defer logger.Sync()

	psqlDB := db.NewPostgre()
	defer db.CloseDB(psqlDB.DB)

	// Initialize repository and service with interfaces
	userRepository := repository.NewUserRepository(psqlDB)
	userService := service.NewUserService(userRepository)

	// Initialize controllers with interfaces
	userController := controller.NewUserController(userService)

	app := fiber.New()

	//TODO: logger middleware will be added later
	// // Middleware
	// app.Use(logger)

	r := router.New(app, userController)
	r.Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
