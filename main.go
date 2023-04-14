package main

import (
	"blog_app_server/handlers/post"
	"blog_app_server/handlers/comment"
	"blog_app_server/db"
	"log"
	"github.com/gin-gonic/gin"
)



// go.mongodb.org/mongo-driver/mongo


func init(){

}



func main(){
	router := gin.Default()
	dbErr:= db.ConnectToDB()
	if dbErr !=nil{
		log.Fatal("Error connecting to database:",dbErr)
	}
		//Post
        router.POST("/create_post",post.CreatePostHandler)
        router.POST("/get_posts",post.GetPostsHandler)
        router.POST("/get_single_post",post.GetSinglePostHandler)
        router.POST("/edit_post",post.EditPostHandler)
        router.POST("/delete_post",post.DeletePostHandler)


		//Comment
		router.POST("/create_comment",comment.CreateCommentHandler)
        router.POST("/get_comments",comment.GetCommentsHandler)
        router.POST("/delete_comment",comment.DeleteCommentHandler)

        err := router.Run("localhost:8080")

	if err != nil{
		log.Fatal("Error starting server:",err)
	}
}