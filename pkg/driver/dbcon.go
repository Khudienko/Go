package driver

import (
	"database/sql"
	"log"
)

var Database *sql.DB

func OpenConnection() {
	db, err := sql.Open("mysql", "root:@tcp(:3306)/productdb")

	if err != nil {
		log.Println(err)
	}
	Database = db
}
