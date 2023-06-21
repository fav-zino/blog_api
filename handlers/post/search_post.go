package post

import (
	"blog_app_server/db"
	"blog_app_server/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func SearchPostsHandler(c *gin.Context){
	var requestBody struct {
		Title string `json:"title"`//required 
	}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": err})
		return
	}
	

	
	filter := bson.M{
		"title": primitive.Regex{Pattern:requestBody.Title,Options:""},
	}
	var posts []model.Post
    cursor,err := db.PostCollection.Find(context.Background(), filter)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
    }
	if err = cursor.All(context.Background(), &posts); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	if len(posts) == 0{
		c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"No post found","count":len(posts)})
		return 
	}
	
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"success","posts": posts,"count":len(posts)})

}