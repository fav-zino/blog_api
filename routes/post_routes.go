package routes

import (
	"blog_app_server/handlers/post"

	"github.com/gin-gonic/gin"
)

func LoadPostRoutes(router *gin.Engine) {

	postRouter := router.Group("/post")
	postRouter.POST("/create_post", post.CreatePostHandler)
	postRouter.POST("/get_posts", post.GetPostsHandler)
	postRouter.POST("/get_single_post", post.GetSinglePostHandler)
	postRouter.POST("/edit_post", post.EditPostHandler)
	postRouter.POST("/delete_post", post.DeletePostHandler)

}
