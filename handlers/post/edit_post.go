package post

import (
	"blog_app_server/db"
	model "blog_app_server/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EditPostHandler(c *gin.Context) {
	var requestBody struct {
		ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`//required
		Title   string             `json:"title"`//required
		Content string             `json:"content"`//required
	}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": err})
		return
	}

	if requestBody.ID == primitive.NilObjectID  {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing required fields: '_id'"})
		return
	}
	if  requestBody.Content == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing required fields: 'content'"})
		return
	}
	if  requestBody.Title == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing required fields: 'title'"})
		return
	}

	timestamp := time.Now().Unix()
	filter := bson.M{"_id": requestBody.ID}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	update := bson.M{"$set": bson.M{"content": requestBody.Content,"title": requestBody.Title, "timestamp": timestamp,}}

	var post model.Post
	err = db.PostCollection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&post)
	if err == mongo.ErrNoDocuments {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Post with this id not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "message": "Post edit successful", "post": post})

}
