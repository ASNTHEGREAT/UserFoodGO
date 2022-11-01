package main

import (
	"net/http"
	"os"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/ASNTHEGREAT/entity"
	"github.com/ASNTHEGREAT/middleware"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger(),
		middleware.BasicAuth())


	userRoutes := server.Group("/users/")
	{
		userRoutes.GET("/getall", func (ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, entity.users)
		})

		userRoutes.POST("/new", func (ctx *gin.Context)  {
			ctx.IndentedJSON(http.StatusOK, entity.AddUser())
		})
	}


	server.Run(":8080")
}