package routes

import (
	"blog_app_server/handlers/comment"
	"github.com/gin-gonic/gin"
)

func LoadCommentRoutes(router *gin.Engine) {

	commentRouter := router.Group("/comment")
	commentRouter.POST("/create_comment",comment.CreateCommentHandler)
	commentRouter.POST("/get_comments",comment.GetCommentsHandler)
	commentRouter.POST("/delete_comment",comment.DeleteCommentHandler)

}
