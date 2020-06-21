package main

import (
	"context"
	"errors"
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
