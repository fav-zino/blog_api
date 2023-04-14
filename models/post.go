package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Post struct{
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Author    string             `bson:"author"`
    Title     string             `bson:"title"`
    Content   string             `bson:"content"`
    ViewCount   int              `bson:"view_count"`
    CommentCount   int            `bson:"comment_count"`
    Timestamp int64        `bson:"timestamp,omitempty"`
}



// func (p *Post) Validate() error{
// 	if p.Body == "" || p.Author == "" {
// 		return errors.New("all required fields must be filled")
// 	}

// 	return nil
// }


// func (p *Post)  Save() error{
// 	_,err := db.PostCollection.InsertOne(context.Background(),p)
// 	if err != nil{
// 		return err
// 	}
// 	return nil
// }