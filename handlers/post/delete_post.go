package post

import (
	"blog_app_server/db"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeletePostHandler(c *gin.Context) {
	var requestBody struct {
		ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`//required
	}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": err})
		return
	}

	if requestBody.ID == primitive.NilObjectID {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing required field: '_id'"})
		return
	}

	// Convert the ID string to a primitive.ObjectID value
	// id, _ := primitive.ObjectIDFromHex( request.ID)
	// bson.ObjectIdHex(postID)

	filter := bson.M{"_id": requestBody.ID}
	res, err := db.PostCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err})
		return
	}

	if res.DeletedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Post with this id not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "message": "Post deleted successful"})

}
