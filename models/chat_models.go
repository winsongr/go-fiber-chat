package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateFetchReq struct {
	UserId primitive.ObjectID `json:"userId"`
	SeconfdUserId primitive.ObjectID `json:"secondUserId"`
}

type CreateFetchRes struct {
	UserId primitive.ObjectID `json:"userId"`
	SeconfdUserId primitive.ObjectID `json:"secondUserId"`
	Users []primitive.ObjectID `json:"users"`
	
}