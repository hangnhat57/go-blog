package global

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

var NilUser User

func (u User) GetToken() string {
	byteSlc, _ := json.Marshal(u)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"data": string(byteSlc),
	})
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}
