package post

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


func EditPostHandler(c *gin.Context){
	var post model.Post
	err:= c.BindJSON(&post); if err !=nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "some error occured"})
		return
	}

	if post.ID == primitive.NilObjectID || post.Content == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "all required fields must be filled"})
		return
	}

	timestamp := time.Now().Unix()
	filter := bson.M{"_id": post.ID}
	update := bson.M{"$set": bson.M{"content": post.Content,"timestamp":timestamp}}
	res, err := db.PostCollection.UpdateOne(context.Background(), filter, update); if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	if res.MatchedCount  == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status":"error","message":"post with this id not found"})
		return
	}

	
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"edit successful","post":post})

}