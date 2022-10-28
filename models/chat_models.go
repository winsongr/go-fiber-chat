package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateChatReq struct {
	UserId       primitive.ObjectID `json:"userId"`
	SecondUserId primitive.ObjectID `json:"secondUserId"`
}

type CreateChatRes struct {
	ChatId          primitive.ObjectID   `json:"chatId"`
	Users           []primitive.ObjectID `json:"users"`
	IsGroup         bool                 `json:"isGroup"`
	LatestMessage   string               `json:"latestMessage"`
	LatestMessageId string               `json:"latestMessageId"`
}

type GetAllChatsReq struct {
	UserId primitive.ObjectID `json:"userId"`
}

type GetAllChatsRes struct {
	UserId primitive.ObjectID `json:"userId"`
	Chats  []CreateChatRes    `json:"chats"`
}

type AddToGroupReq struct {
	UserId primitive.ObjectID `json:"userId"`
	ChatId primitive.ObjectID `json:"chatId"`
}

type DeleteFromGroupReq struct {
	UserId primitive.ObjectID `json:"userId"`
	ChatId primitive.ObjectID `json:"chatId"`
}
