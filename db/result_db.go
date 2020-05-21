package db

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

func dbInit(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error creating database table: %q", err))
            return
        }

        if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error incrementing tick: %q", err))
            return
        }

        rows, err := db.Query("SELECT tick FROM ticks")
        if err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error reading ticks: %q", err))
            return
        }

        defer rows.Close()
        for rows.Next() {
            var tick time.Time
            if err := rows.Scan(&tick); err != nil {
                c.String(http.StatusInternalServerError,
                    fmt.Sprintf("Error scanning ticks: %q", err))
                return
            }
            c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
        }
    }
}
