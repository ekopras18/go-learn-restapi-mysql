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

	// authentication route
	authentication := router.Group("api/auth")
	{
		authentication.POST("/login", auth.Login)
		authentication.POST("/register", auth.Register)
		authentication.GET("/logout", auth.Logout)
	}

	// base route
	router.GET("/", base.Index)

	// search route
	search := router.Group("api/search")
	{
		search.GET("/blogs", blogcontroller.Search)
	}

	// api/v1 route without JWT middleware
	noauth := router.Group("api/v1")
	{
		noauth.GET("/blogs", blogcontroller.Index)
		noauth.GET("/blog/:id", blogcontroller.Show)
	}

	// api/v1 route with JWT middleware
	v1 := router.Group("api/v1")
	{
		v1.Use(auth.JWTMiddleware())
		// blog
		v1.POST("/blog", blogcontroller.Create)
		v1.PUT("/blog/:id", blogcontroller.Update)
		v1.DELETE("/blog/:id", blogcontroller.Delete)

		// more controller

	}
	return router
}
