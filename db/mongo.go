package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var PostCollection *mongo.Collection
var CommentCollection *mongo.Collection

func ConnectToDB() error{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client,err := mongo.Connect(context.Background(),clientOptions)
	if err != nil {
		return err
	}
	PostCollection = client.Database("mydb").Collection("post")
	CommentCollection = client.Database("mydb").Collection("comment")
	return nil
	
}