package sql_db

import (
	"database/sql"
	"log"
)

func ConnectToDb(c DbConnection) *sql.DB {
	db, err := sql.Open("sqlserver", c.GetConnectionString())

	if err != nil {
		log.Fatalf("Cannot connect to Server=%v Database=%v: %v\n", c.Server, c.Database, err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot ping Server=%v Database=%v: %v\n", c.Server, c.Database, err)
	}
	return db
}
