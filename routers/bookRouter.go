package router

import (
	"challange-7/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controllers.AddBook)

	router.GET("/book/:ID", controllers.GetBook)

	router.GET("/books", controllers.GetBooks)
	
	router.PUT("/book/:ID", controllers.UpdateBook)

	router.DELETE("/book/:ID", controllers.DeleteBook)

	return router
}