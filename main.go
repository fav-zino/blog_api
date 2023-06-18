package main

import (
	"blog_app_server/db"
	"blog_app_server/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	dbErr := db.ConnectToDB()
	if dbErr != nil {
		log.Fatal("Error connecting to database:", dbErr)
	}
	routes.LoadPostRoutes(router)
	routes.LoadCommentRoutes(router)

	gin.SetMode(gin.DebugMode)
	err := router.Run("localhost:8080")

	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
