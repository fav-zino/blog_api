package comment

import (
	"blog_app_server/db"
	"blog_app_server/models"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func GetCommentsHandler(c *gin.Context){
	var requestBody model.Comment
	err:= c.BindJSON(&requestBody); if err !=nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": err})
		return
	}

	
	if requestBody.PostID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status":"error","message": "Missing required field: 'post_id'"})
		return
	}

	var comments []model.Comment
	filter := bson.M{"postID": requestBody.PostID}
    cursor,err := db.CommentCollection.Find(context.Background(), filter)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
    }
	if err = cursor.All(context.Background(), &comments); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	if len(comments) == 0{
		c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"No comment at the moment","count":len(comments)})
		return 
	}
	
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"success","comments": comments,"count":len(comments)})

}