package main

import (
    "log"
    "os"
    "database/sql"
    "net/http"
    "github.com/cenkalti/backoff/v4"
    "github.com/go-sql-driver/mysql"
    "github.com/gin-gonic/gin"
)

var (
    db *sql.DB
)

type Book struct {
    BookID uint64 `json:"book_id"`
    Title string `json:"title"`
    Author string `json:"author"`
}

func main() {
    config := mysql.Config{
        User: os.Getenv("MARIADB_USER"),
        Passwd: os.Getenv("MARIADB_PASSWORD"),
        Net: "tcp",
        Addr: os.Getenv("MYSQL_ADDR"),
        DBName: os.Getenv("APP_DB_NAME"),
        AllowNativePasswords: true,
    }
    var err error
    openDB := func() error {
        db, err = sql.Open("mysql", config.FormatDSN())
        return err
    }
    err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    tryPing := func() error {
        err = db.Ping()
        return err
    }
    err = backoff.Retry(tryPing, backoff.NewExponentialBackOff())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }

    router := gin.Default()
    router.GET("/books", GetAllBooks)
    router.Run()
}

func GetAllBooks(ctx *gin.Context) {
    rows, err := db.Query("SELECT * FROM books")
    log.Println("GetAllBooks")
    if err != nil {
        ctx.String(http.StatusInternalServerError, "")
        return
    }
    defer rows.Close()
    var books []Book
    for rows.Next() {
        var book Book
        if err := rows.Scan(&book.BookID, &book.Title, &book.Author); err != nil {
            ctx.String(http.StatusInternalServerError, "")
            return
        }
        books = append(books, book)
    }
    ctx.JSON(http.StatusOK, books)
}
