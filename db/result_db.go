package main

import (
	"bytes"
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func dbFunc(db *sql.DB) gin.HandlerFunc {

}

func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
	db, err := sql.Open("mysql", os.Getenv("CLEARDB_DATABASE_URL"))
	

    // if there is an error opening the connection, handle it
    if err != nil {
		log.Fatalf("Error opening database: %q", err)
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("public ")


}