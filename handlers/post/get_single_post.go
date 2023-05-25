package post

import (
	"blog_app_server/db"
	"blog_app_server/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func GetSinglePostHandler(c *gin.Context){
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

	
	update := bson.M{
		"$inc": bson.M{"view_count": 1},
    }
	filter := bson.M{"_id":  request.ID }
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var result model.Post
	if err = db.PostCollection.FindOneAndUpdate(context.Background(), filter,update,options).Decode(&result); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"success","post": result})

}