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
	initANewUser()
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

func Test_authServer_SignUp(t *testing.T) {
	initANewUser()
	server := authServer{}
	_, err := server.SignUp(context.Background(), &proto.SignUpRequest{
		Username: "Nathan",
		Email:    "kathyhoang1911@gmail.com",
		Password: "convitconga",
	})
	if err == nil {
		t.Error("Didnot validate Username")
	}
	_, err = server.SignUp(context.Background(), &proto.SignUpRequest{
		Username: "Kathy",
		Email:    "hangnhat57@gmail.com",
		Password: "1concho1conmeo",
	})
	if err == nil {
		t.Error("Didnot validate email")
	}
	_, err = server.SignUp(context.Background(), &proto.SignUpRequest{
		Username: "Kathy",
		Email:    "kathyhoang1911@gmail.com",
		Password: "1concho1conmeo",
	})
	if err != nil {
		t.Error("Something wrong, could not create new user", err)
	}
}

func initANewUser() {
	global.ConnectToTestDB()
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	global.DB.Collection("user").InsertOne(context.Background(), global.User{
		ID:       primitive.NewObjectID(),
		Email:    "hangnhat57@gmail.com",
		Username: "Nathan",
		Password: string(pw),
	})
}
