package main

import (
	"context"
	"go-blog/backend/global"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	global.DB.Collection("user").InsertOne(context.Background(), bson.M{"name": "test"})
}
