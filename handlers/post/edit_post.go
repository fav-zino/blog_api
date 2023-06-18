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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func EditPostHandler(c *gin.Context){
	var  requestBody struct{
		ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id" binding:"required"`
		Title     string             `json:"title"`
		Content   string             `json:"content"`
	}
	err:= c.BindJSON(&requestBody); if err !=nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": err})
		return
	}

	if requestBody.ID == primitive.NilObjectID &&  requestBody.Content == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "Missing required fields: '_id','content'"})
		return
	}


	timestamp := time.Now().Unix()
	filter := bson.M{"_id": requestBody.ID}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	update := bson.M{"$set": bson.M{"content": requestBody.Content,"timestamp":timestamp}}
	
	var post model.Post
	err = db.PostCollection.FindOneAndUpdate(context.Background(), filter, update,opts).Decode(&post)	
	if err == mongo.ErrNoDocuments {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Post with this id not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}
	
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"Post edit successful","post":post})

}