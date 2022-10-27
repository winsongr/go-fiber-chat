package controllers

import (
	"github.com/achintya-7/go-fiber-chat/configs"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var chatCollection *mongo.Collection = configs.GetCollection(configs.DB, "chats")

func ChatRoute(app *fiber.App) {
	
}