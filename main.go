package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamalogudah/booky/controllers"
	"github.com/kamalogudah/booky/models"
)

func main() {
	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	//   c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.Run()
}
