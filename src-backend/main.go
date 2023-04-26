package main

import (
	"TUBES3_13521102/src-backend/controller"
	"TUBES3_13521102/src-backend/database"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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


var (
  db *sql.DB
)


func main() {
    // date := "25/02/2023"

    // dayName := FeatureDate.DateDayName(date)
    // fmt.Println(dayName)

    // hasil, err := FeatureCalculator.CalculateExpression("((-8.05) + 3)")

    // if (err != nil) {
    //   fmt.Println(err)
    // } else {
    //   fmt.Println(hasil)
    // }

    // tes, err := repository.GetHistoryByHistoryID(db, 1)

    // if err != nil {
    //   fmt.Println(err.Error())
    // }

    // fmt.Println(tes.HistoryID)
    // fmt.Println(tes.Topic)

    // for _, message := range tes.Conversation {
    //   fmt.Println(message.Id)
    //   fmt.Println(message.Text)
    //   fmt.Println(message.SentTime)
    //   fmt.Println(message.HistoryId)
    // }

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

    r := gin.Default()
    
    // CORS middleware
    r.Use(func(c *gin.Context) {
      c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
      c.Next()
    })

    r.GET("/history", controller.GetAllHistoryMessage)

    r.POST("/message", controller.InsertUserMessage)

    r.Run()
}
