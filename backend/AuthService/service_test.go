package main

import (
	"context"
	"go-blog/backend/global"
	"go-blog/backend/proto"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Test_authServer_Login(t *testing.T) {
	global.ConnectToTestDB()
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	global.DB.Collection("user").InsertOne(context.Background(), global.User{
		ID:       primitive.NewObjectID(),
		Email:    "hangnhat57@gmail.com",
		Username: "Nathan",
		Password: string(pw),
	})
	server := authServer{}
	_, err := server.Login(context.Background(), &proto.LogInRequest{
		Login:    "hangnhat57@gmail.com",
		Password: "password",
	})
	if err != nil {
		t.Error("Error:", err.Error())
	}
	_, err = server.Login(context.Background(), &proto.LogInRequest{
		Login:    "whatever@whicheve.com",
		Password: "Pass ne",
	})
	if err == nil {
		t.Error("Error was nil")
	}
}
