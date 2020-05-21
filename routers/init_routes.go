package routers

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func initialize_routes() {	
	
	db, err := sql.Open("mysql", os.Getenv("CLEARDB_DATABASE_URL"))
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()


	router.GET("/", show_index_page)
	//router.GET("/statistics", showStats(db))

	// TODO: fill out the rest of routes
}
