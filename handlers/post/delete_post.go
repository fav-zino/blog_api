package post

import (
	"blog_app_server/db"
	model "blog_app_server/models"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)




func DeletePostHandler(c *gin.Context){
	var request model.Post
	err:= c.BindJSON(&request); if err !=nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "some error occured"})
		return
	}

	
	if request.ID == primitive.NilObjectID {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "all required fields must be filled"})
		return
	}

	// Convert the ID string to a primitive.ObjectID value
	// id, _ := primitive.ObjectIDFromHex( request.ID)
	// bson.ObjectIdHex(postID)
	
	filter := bson.M{"_id": request.ID}
	res,err := db.PostCollection.DeleteOne(context.Background(), filter);if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	if res.DeletedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status":"error","message":"post with this id not found"})
		return
	}

	
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"delete successful"})

}