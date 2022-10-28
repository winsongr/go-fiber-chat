package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/achintya-7/go-fiber-chat/configs"
	"github.com/achintya-7/go-fiber-chat/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var chatCollection *mongo.Collection = configs.GetCollection(configs.DB, "chats")

func GetFetchChats(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var chat models.CreateFetchReq

	if err := c.BodyParser(&chat); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": err,
		})	
	}

	chatFilter := bson.D{{Key: "isGroup", Value: false}, [{}, {}]}
	isChat := chatCollection.FindOne(ctx, chatFilter)


}