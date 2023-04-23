package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Variables to connect to database host

func ConnectDatabase (username string, password string, host string, port string, database string) string {
  // Testing connecct to database
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
}
