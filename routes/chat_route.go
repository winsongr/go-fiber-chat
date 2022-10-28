package routes

import (
	"github.com/achintya-7/go-fiber-chat/controllers"
	"github.com/gofiber/fiber/v2"
)

func ChatRoute(app *fiber.App) {
	app.Post("/create_fetch", controllers.GetFetchChats)
}