package repository

import (
	"database/sql"
	"log"
	"os"

	// Import MYSQL Driver
	_ "github.com/go-sql-driver/mysql"
	// Import PostgreSQL Driver
	_ "github.com/lib/pq"
	// Import SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *sql.DB {
	db, err := sql.Open(os.Getenv("GOSCOPE_DATABASE_TYPE"), os.Getenv("GOSCOPE_DATABASE_CONNECTION"))
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	return db
}
