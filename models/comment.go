package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct{
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Author    string             `json:"author"`
    Content   string             `json:"content"`
    Timestamp int64             `json:"timestamp,omitempty"`
	PostID    string            `json:"post_id"`
}