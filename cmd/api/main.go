package main

import (
    "github.com/gin-gonic/gin"
    "FirstApi/api"
)

func main() {
    router := gin.Default()
    // Adding the routers
    router.GET("/books", api.GetAllBooks) // returns a JSON with all books
    router.GET("/books/:id", api.GetBookById)
    router.POST("/books", api.AddBook)
    router.Run("localhost:8090")
}
