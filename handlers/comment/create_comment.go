package comment

import (
	"blog_app_server/db"
	model "blog_app_server/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCommentHandler(c *gin.Context) {
	var requestBody struct {
		Author  string `json:"author"`//required
		Content string `json:"content"`//required
		PostID  string `json:"post_id"`//required
	}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": err})
		return
	}

	if requestBody.Author == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing required field: 'author'"})
		return
	}
	if requestBody.Content == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing required field: 'content'"})
		return
	}
	if requestBody.PostID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing required field: 'post_id'"})
		return
	}

	timestamp := time.Now().Unix()
	var comment model.Comment
	comment.Author = requestBody.Author
	comment.Content = requestBody.Content
	comment.Timestamp = timestamp
	comment.PostID = requestBody.PostID

	result, err := db.CommentCollection.InsertOne(context.Background(), comment)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}
	comment.ID = result.InsertedID.(primitive.ObjectID)

	//increment comment count for post
	postID, _ := primitive.ObjectIDFromHex(comment.PostID)
	filter := bson.M{"_id": postID}
	update := bson.M{
		"$inc": bson.M{"comment_count": 1},
	}
	res, err := db.PostCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}

	if res.MatchedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Post with this id not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "message": "success", "comment": comment})

}
