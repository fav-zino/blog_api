package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct{
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Author    string             `bson:"author"`
    Title     string             `bson:"title"`
    Content   string             `bson:"content"`
    Timestamp int64             `bson:"timestamp,omitempty"`
	PostID    string            `bson:"postID"`
}