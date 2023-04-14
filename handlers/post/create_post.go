package post

import (
	"blog_app_server/db"
	"blog_app_server/models"
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreatePostHandler(c *gin.Context){
	var post model.Post
	err:= c.BindJSON(&post); if err !=nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "some error occured"})
		return
	}

	if post.Author  == "" || post.Content == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "all required fields must be filled"})
		return
	}

	timestamp := time.Now().Unix()
	post.Timestamp = timestamp

	result,err := db.PostCollection.InsertOne(context.Background(),post);if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	post.ID = result.InsertedID.(primitive.ObjectID)
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"success","post": post})

}