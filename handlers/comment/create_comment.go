package comment

import (
	"blog_app_server/db"
	"blog_app_server/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateCommentHandler(c *gin.Context){
	var comment model.Comment
	err:= c.BindJSON(&comment); if err !=nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "some error occured"})
		return
	}

	if comment.Author  == "" || comment.Content == "" || comment.PostID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "all required fields must be filled"})
		return
	}

	timestamp := time.Now().Unix()
	comment.Timestamp = timestamp

	result,err := db.CommentCollection.InsertOne(context.Background(),comment);if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}
	comment.ID = result.InsertedID.(primitive.ObjectID)

	//increment comment count for post
	postID, _ := primitive.ObjectIDFromHex( comment.PostID)
	filter := bson.M{"_id": postID }
	update := bson.M{
		"$inc": bson.M{"comment_count": 1},
    }
	res, err := db.PostCollection.UpdateOne(context.Background(), filter, update); if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	if res.MatchedCount  == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status":"error","message":"post with this id not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"success","comment": comment})

}