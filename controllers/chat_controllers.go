package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/achintya-7/go-fiber-chat/configs"
	"github.com/achintya-7/go-fiber-chat/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var chatCollection *mongo.Collection = configs.GetCollection(configs.DB, "chats")

func CreateChat(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var chat models.CreateChatReq
	var resChat models.CreateChatRes

	if err := c.BodyParser(&chat); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Status":  err,
			"Message": err.Error(),
			"Data":    nil,
		})
	}

	chatFilter := bson.D{{Key: "users", Value: bson.A{chat.UserId, chat.SecondUserId}}}
	isChat := chatCollection.FindOne(ctx, chatFilter)

	// no document was found
	if isChat.Err() != nil {

		chatNew := models.CreateChatRes{
			ChatId:          primitive.NewObjectID(),
			IsGroup:         false,
			Users:           []primitive.ObjectID{chat.UserId, chat.SecondUserId},
			LatestMessage:   "",
			LatestMessageId: "",
		}

		result, err := chatCollection.InsertOne(ctx, chatNew)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"Status":  err,
				"Message": err.Error(),
				"Data":    result,
			})
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Status":  "Suceess",
			"Message": "Chat Created",
			"Data":    chatNew,
		})
	}

	// a chat was found
	isChat.Decode(&resChat)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Status":  "Success",
		"Message": "Chat already present",
		"Data":    resChat,
	})
}

func AddToGroup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var req models.AddToGroupReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Status":  err,
			"Message": err.Error(),
			"Data":    nil,
		})
	}

	filter := bson.D{{Key: "_id", Value: req.ChatId}}
	update := bson.D{
		{
			Key: "$push",
			Value: bson.D{
				{
					Key:   "users",
					Value: req.ChatId,
				},
			},
		},
	}

	err := chatCollection.FindOneAndUpdate(ctx, filter, update)
	if err.Err() != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Status":  err,
			"Message": err.Err(),
			"Data":    "Cant Add User",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"Status":  "Success",
		"Message": "User Added",
		"Data":    err,
	})
}

func DeleteFromGroup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var req models.DeleteFromGroupReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Status":  err,
			"Message": err.Error(),
			"Data":    nil,
		})
	}

	filter := bson.D{{Key: "_id", Value: req.ChatId}}
	update := bson.D{
		{
			Key: "$pull",
			Value: bson.D{
				{
					Key:   "users",
					Value: req.ChatId,
				},
			},
		},
	}

	err := chatCollection.FindOneAndUpdate(ctx, filter, update)
	if err.Err() != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Status":  err,
			"Message": err.Err(),
			"Data":    "Cant Delete User",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"Status":  "Success",
		"Message": "User Removed",
		"Data":    err,
	})
}

func GetAllChats(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var chat models.GetAllChatsReq
	var chats []models.CreateChatRes

	if err := c.BodyParser(&chat); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Status":  err,
			"Message": err.Error(),
			"Data":    "",
		})
	}

	results, err := chatCollection.Find(ctx, bson.D{{Key: "users", Value: chat.UserId}})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Status":  err.Error(),
			"Message": err.Error(),
			"Data":    results,
		})
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleChat models.CreateChatRes
		if err = results.Decode(&singleChat); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"Status":  err.Error(),
				"Message": err.Error(),
				"Data":    "",
			})
		}
		chats = append(chats, singleChat)
	}

	return c.Status(200).JSON(fiber.Map{
		"Status":  "Success",
		"Message": "All Chats provided",
		"Data":    chats,
	})
}
