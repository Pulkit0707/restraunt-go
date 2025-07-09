package database

import(
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance()*mongo.Client{
	MongoDb:="mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(MongoDb)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
    return client
}

var Client*mongo.Client=DBinstance()

func OpenCollection(client*mongo.Client,collectionName string)*mongo.Collection{
	var collection *mongo.Collection=client.Database("restraunt").Collection(collectionName)
	return collection
}

