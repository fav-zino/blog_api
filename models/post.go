package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Post struct{
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Author    string             `json:"author"`
    Title     string             `json:"title"`
    Content   string             `json:"content"`
    ViewCount   int              `json:"view_count"`
    CommentCount   int            `json:"comment_count"`
    Timestamp int64        `json:"timestamp,omitempty"`
}

