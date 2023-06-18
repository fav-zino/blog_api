package post

import (
	"blog_app_server/db"
	"blog_app_server/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func GetSinglePostHandler(c *gin.Context){
	var  requestBody struct{
		ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id" binding:"required"`
	}

	err:= c.BindJSON(&requestBody); if err !=nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": err})
		return
	}

	
	if requestBody.ID == primitive.NilObjectID {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "Missing required field: '_id'"})
		return
	}

	// Convert the ID string to a primitive.ObjectID value
	// id, _ := primitive.ObjectIDFromHex( request.ID)

	
	update := bson.M{
		"$inc": bson.M{"view_count": 1},
    }
	filter := bson.M{"_id":  requestBody.ID }
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var result model.Post
	if err = db.PostCollection.FindOneAndUpdate(context.Background(), filter,update,options).Decode(&result); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	if err == mongo.ErrNoDocuments {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Post with this id not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}

	
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"success","post": result})

}