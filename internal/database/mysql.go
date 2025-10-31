package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var MySQLDB *sql.DB

// InitMySQL initializes the MySQL database connection
func InitMySQL() error {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	// Create DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, database)

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening MySQL connection: %w", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error pinging MySQL database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	MySQLDB = db
	log.Println("✅ MySQL connected successfully")
	return nil
}

// CloseMySQL closes the MySQL database connection
func CloseMySQL() {
	if MySQLDB != nil {
		MySQLDB.Close()
		log.Println("MySQL connection closed")
	}
}
