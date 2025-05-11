package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "FirstApi/datas"
    "FirstApi/dto"
    "strconv"
)

func GetAllBooks(c *gin.Context) {
    books, err := datas.GetAllBooks()
    if err != nil {
        return
    }
    c.IndentedJSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
    var book_request dto.BookRequest

    err := c.BindJSON(&book_request)
    if err != nil {
        c.JSON(http.StatusBadRequest, err.Error())
        return
    }
    new_book := datas.AddBook(book_request)
    c.IndentedJSON(http.StatusCreated, new_book)
}

func GetBookById(c *gin.Context) {
    book_id, err1 := strconv.Atoi(c.Param("id"))

    if err1 != nil {
        c.JSON(http.StatusBadRequest, err1.Error())
        return
    }

    my_book , err := datas.GetBookById(book_id)

    if err != nil {
        c.JSON(http.StatusBadRequest, err.Error())
        return
    }

    c.IndentedJSON(http.StatusOK, my_book)
}
