package database

import (
	"daemon_backend.bin/component/extractor"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func dbURLGenerator() string {
	dbuser, dbpassword := extractor.ExtractStrFromFile("database", "user"), extractor.ExtractStrFromFile("database", "password")
	dbname, dbaddress := extractor.ExtractStrFromFile("database", "name"), extractor.ExtractStrFromFile("database", "address")

	dbURL := dbuser + ":" + dbpassword + "@tcp(" + dbaddress + ")/" + dbname

	return dbURL
}

func DbConnector() (*sql.DB, error) {
	dbengine := extractor.ExtractStrFromFile("database", "engine")
	dbURL := dbURLGenerator()

	conn, err := sql.Open(dbengine, dbURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return conn, nil
}
