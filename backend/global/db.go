package global

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database

func connectToDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburl))
	if err != nil {
		log.Fatal(err)
	}
	DB = *client.Database(dbname)
}

//NewDBContext return a new Context following app performance
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*performance/100)
}

//ConnectToTestDB overwrite real DB for testing
func ConnectToTestDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburl))
	if err != nil {
		log.Fatal(err)
	}
	DB = *client.Database(dbname + "_test")
}
