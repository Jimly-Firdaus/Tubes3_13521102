package main

import (
    "database/sql"
    "fmt"
    "log"
    "TUBES3_13521102/src-backend/featureDate"
    "TUBES3_13521102/src-backend/featureCalculator"
    "TUBES3_13521102/src-backend/database"
    _ "github.com/go-sql-driver/mysql"
)

// Variables to connect to database host
const (
  username = "root"
  host = "containers-us-west-13.railway.app"
  password = "PNGO6atNekbjjq4g2yPy"
  port = "6330"
  databasetype = "railway"
)

func main() {
    date := "25/02/2023"

    dayName := FeatureDate.DateDayName(date)
    fmt.Println(dayName)

    hasil, err := FeatureCalculator.CalculateExpression("8 + 7 + 2+* 5")

    if (err != nil) {
      fmt.Println(err)
    } else {
      fmt.Println(hasil)
    }

    db, err := sql.Open("mysql", database.ConnectDatabase(username, password, host, port, databasetype))
    if (err != nil) {
      fmt.Printf("Error %s while opening database\n", err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
    defer db.Close()
}
