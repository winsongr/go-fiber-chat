package controllers

import (
	"github.com/achintya-7/go-fiber-chat/configs"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func UserRoute(app *fiber.App) {

}
