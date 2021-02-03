package testutils

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/db"
)

const (
	dbHost     = "mysql"
	dbPort     = 3306
	dbUser     = "root"
	dbPassword = ""
	dbName     = "go_http_server_sample_test"
)

var (
	dbConn *sql.DB
)

// InitDB connects the DB and returns the function to close the connection.
func InitDB() func() {
	conf := db.Config{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
	}

	db, err := sql.Open("mysql", conf.DSN())
	if err != nil {
		log.Fatalf("failed to connect to the DB: %s", err)
	}

	dbConn = db

	return func() {
		dbConn.Close()
	}
}

// SetUpDB inserts test data to the DB. It returns the DB connection and the function to truncate tables.
func SetUpDB() (*sql.DB, func()) {
	if dbConn == nil {
		log.Fatal("DB connection not initialized")
	}

	ctx := context.Background()

	setUpFixtures(ctx)

	return dbConn, func() {
		truncateTables(ctx)
	}
}

func createTablesIfNotExist(ctx context.Context) {
	executeSQLScript(ctx, "../../configs/test/sql/schema.sql")
}

func setUpFixtures(ctx context.Context) {
	executeSQLScript(ctx, "../../configs/test/sql/fixture.sql")
}

func executeSQLScript(ctx context.Context, path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read %s: %v", path, err)
	}

	queries := strings.Split(string(b), ";")

	for _, query := range queries {
		if len(strings.TrimSpace(query)) == 0 {
			continue
		}
		if _, err := dbConn.ExecContext(ctx, query); err != nil {
			log.Fatalf("failed to execute query: %v", err)
		}
	}
}

func truncateTables(ctx context.Context) {
	rows, err := dbConn.QueryContext(ctx, "SHOW TABLES")
	if err != nil {
		log.Fatalf("failed to show tables: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("failed to scan table: %v", err)
		}
		if _, err := dbConn.ExecContext(ctx, fmt.Sprintf("TRUNCATE %s", tableName)); err != nil {
			log.Fatalf("failed to truncate %s: %v", tableName, err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("failed to scan tables: %v", err)
	}
}
