package main

import (
	"context"
	"errors"
	"fmt"
	"go-blog/backend/global"
	"go-blog/backend/proto"
	"log"
	"net"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type authServer struct{}

func (authServer) Login(_ context.Context, in *proto.LogInRequest) (*proto.AuthResponse, error) {

	login, password := in.GetLogin(), in.GetPassword()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var user global.User
	global.DB.Collection("user").FindOne(ctx, bson.M{"$or": []bson.M{bson.M{"username": login}, bson.M{"email": login}}}).Decode(&user)
	if user == global.NilUser {
		return &proto.AuthResponse{}, errors.New("Invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return &proto.AuthResponse{}, errors.New("Invalid Password")
	}
	return &proto.AuthResponse{Token: user.GetToken()}, nil
}

func (authServer) SignUp(_ context.Context, in *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	userName, email, password := in.GetUsername(), in.GetEmail(), in.GetPassword()

	if len(userName) < 3 || len(email) < 3 || len(password) < 3 {
		return nil, errors.New("Invalid input")
	}
	if !isValidEmail(email) {
		return nil, errors.New("Duplicated email")
	}
	if !isValidUserName(userName) {
		return nil, errors.New("Duplicated userName")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := global.User{
		Username: userName,
		Email:    email,
		Password: string(hashedPassword),
	}
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	global.DB.Collection("user").InsertOne(ctx, &user)
	return &proto.SignUpResponse{Msg: "Success"}, nil
}
func isValidUserName(username string) bool {
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var user global.User
	global.DB.Collection("user").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if user != global.NilUser {
		return false
	}
	return true

}

func isValidEmail(email string) bool {
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var user global.User
	global.DB.Collection("user").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	fmt.Printf(user.Email)
	if user != global.NilUser {
		return false
	}
	return true

}
func main() {
	gs := grpc.NewServer()
	proto.RegisterAuthServiceServer(gs, authServer{})
	listener, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatal("Port exists", err)
		os.Exit(1)
	}
	gs.Serve(listener)
}
