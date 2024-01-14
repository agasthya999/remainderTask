package router

import (
	handler "remainderTask/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	userApi := api.Group("/user")
	taskApi := api.Group("/task")
	userApi.Post("/addUser", handler.AddUser)
	userApi.Get("/getUser", handler.GetAllUsers)
	taskApi.Post("addTask", handler.AddTask)
	userApi.Get("/addContact", handler.AddContactToUser)
}
