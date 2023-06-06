package main

import (
	"go-learn-restapi-mysql/config"
	"go-learn-restapi-mysql/controllers/auth"
	"go-learn-restapi-mysql/controllers/base"
	"go-learn-restapi-mysql/controllers/blogcontroller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := SetupRouter()

	log.Println("Server started on: http://localhost:8888")
	err := r.Run(":8888")
	if err != nil {
		panic(err)
	}

}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", base.Index)

	au := router.Group("api/auth")
	{
		au.POST("/login", auth.Login)
		au.POST("/register", auth.Register)
		au.GET("/logout", auth.Logout)
	}

	v1 := router.Group("api/v1")
	{
		v1.Use(auth.JWTMiddleware())
		// blog
		v1.GET("/blog", blogcontroller.Index)
		v1.POST("/blog", blogcontroller.Create)
		v1.GET("/blog/:id", blogcontroller.Show)
		v1.PUT("/blog/:id", blogcontroller.Update)
		v1.DELETE("/blog/:id", blogcontroller.Delete)

		// more controller
	}
	return router
}
