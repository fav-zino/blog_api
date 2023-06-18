package comment

import (
	"blog_app_server/db"
	model "blog_app_server/models"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)




func DeleteCommentHandler(c *gin.Context){
	var requestBody model.Comment
	err:= c.BindJSON(&requestBody); if err !=nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": err})
		return
	}

	
	if requestBody.ID == primitive.NilObjectID  {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "Missing required field: '_id'"})
		return
	}
	if  requestBody.PostID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "Missing required field: 'post_id'"})
		return
	}

	filter := bson.M{"_id": requestBody.ID}
	res,err := db.CommentCollection.DeleteOne(context.Background(), filter);if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	if res.DeletedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status":"error","message":"Comment with this id not found"})
		return
	}

	//decrement comment count
	postID, _ := primitive.ObjectIDFromHex( requestBody.PostID)
	filter = bson.M{"_id": postID }
	update := bson.M{
		"$inc": bson.M{"comment_count": -1},
    }
	postRes, err := db.PostCollection.UpdateOne(context.Background(), filter, update); if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	if postRes.MatchedCount  == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status":"error","message":"Post with this id not found"})
		return
	}

	
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"Delete successful"})

}