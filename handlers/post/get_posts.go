package post

import (
	"blog_app_server/db"
	"blog_app_server/models"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func GetPostsHandler(c *gin.Context){

	var posts []model.Post
    cursor,err := db.PostCollection.Find(context.Background(), bson.M{})
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
    }
	if err = cursor.All(context.Background(), &posts); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status":"error","message": err})
		return
	}

	if len(posts) == 0{
		c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"No post at the moment","count":len(posts)})
		return 
	}
	
	c.IndentedJSON(http.StatusOK, gin.H{"status":"ok","message":"success","posts": posts,"count":len(posts)})

}