package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "host=poojanotificationdb29.postgres.database.azure.com port=5432 user=postgresadmin password=Poojasachin@29 dbname=notificationdb sslmode=require"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to Azure PostgreSQL")
	return db, nil
}
