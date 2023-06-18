package post

import (
	"blog_app_server/db"
	model "blog_app_server/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreatePostHandler(c *gin.Context){
	
var  requestBody struct{
	Author    string             `json:"author" binding:"required"`
    Title     string             `json:"title"`
    Content   string             `json:"content" binding:"required"`
}

	err:= c.BindJSON(&requestBody); if err !=nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": err})
		return
	}

	if requestBody.Author  == ""  {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "Missing required field: 'author'"})
		return
	}
	
	if requestBody.Content == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message":"Missing required field: 'content'"})
		return
	}

	timestamp := time.Now().Unix()
	var post model.Post
	post.Author = requestBody.Author
	post.Content = requestBody.Content
	post.Title = requestBody.Title	
	post.Timestamp = timestamp

	result,err := db.PostCollection.InsertOne(context.Background(),post);if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	post.ID = result.InsertedID.(primitive.ObjectID)
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"success","post": post})

}