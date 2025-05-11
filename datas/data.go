package datas

import (
    "errors"
    "database/sql"
    "FirstApi/dto"
    // Driver For Sqlite3 Database
    _ "github.com/mattn/go-sqlite3"
    "log"
    "fmt"
)

type Book struct {
    Id int32 `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
    Quantity int32 `json:"quantity"`
}


func GetBookById(id int) (Book, error) {
    db, err := sql.Open("sqlite3", "./datas/books_database.db")
    treatError(err)

    resp, err := db.Query("SELECT * FROM Books WHERE Id=?", id)
    treatError(err)

    var book Book
    if !resp.Next() {
        return book, errors.New("Id not found")
    }

    resp.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
    return book, nil
}

func GetAllBooks() ([]Book, error){
    var books []Book
    db, err := sql.Open("sqlite3", "./datas/books_database.db")
    treatError(err)

    resp, err := db.Query("SELECT * FROM Books")
    treatError(err)

    for resp.Next() {
        var book Book
        resp.Scan(&book.Id, &book.Title, &book.Author, &book.Quantity)
        books = append(books, book)
    }
    return books, nil
}

func AddBook(book dto.BookRequest) (Book) {
    var new_book Book

    db, err := sql.Open("sqlite3", "./datas/books_database.db")
    treatError(err)

    // Check if the book already exists in database
    resp, err := db.Query("SELECT * FROM Books WHERE Title=? AND Author=?", book.Title, book.Author)

    treatError(err)

    if resp.Next() {
        resp.Scan(&new_book.Id, &new_book.Title, &new_book.Author, &new_book.Quantity)
        resp.Close()

        new_quantity := new_book.Quantity+1

        stmt_connection, err := sql.Open("sqlite3", "./datas/books_database.db")
        treatError(err)

        statement, err := stmt_connection.Prepare("UPDATE Books SET Quantity=? WHERE Id=?")
        treatError(err)

        res, err := statement.Exec(new_quantity, new_book.Id)
        treatError(err)

        fmt.Println(res.RowsAffected())
        stmt_connection.Close()
        return new_book
    }

    _, err = db.Exec("INSERT INTO Books VALUES(NULL, ?, ?, 1)", book.Title, book.Author)
    treatError(err)
    db.Close()
    return new_book
}

func treatError(err error) {
    if (err != nil) {
        log.Fatal(err.Error())
    }
}
