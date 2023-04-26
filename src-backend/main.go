package main

import (
	"TUBES3_13521102/src-backend/database"
	"TUBES3_13521102/src-backend/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"

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

func main() {

  r := gin.Default()
  // CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})
  r.GET("/history", func(c *gin.Context) {
    db, err := sql.Open("mysql", database.ConnectDatabase(username, password, host, port, databasetype))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
    defer db.Close()

    historyList, err := repository.GetAllHistory(db)
    c.JSON(http.StatusOK, gin.H{
        "history": historyList,
    })
  })

  r.GET("/botresponse", func(c *gin.Context) {
    db, err := sql.Open("mysql", database.ConnectDatabase(username, password, host, port, databasetype))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    defer db.Close()

    responseList, err := repository.GetAllBotResponse(db)

    c.JSON(http.StatusOK, gin.H{
        "botResponse": responseList,
    })
  })
  r.Run(); // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

  // date := "25/02/2023"

  //   dayName := FeatureDate.DateDayName(date)
  //   fmt.Println(dayName)

  //   hasil, err := FeatureCalculator.CalculateExpression("((-80.05) / (-1))")

  //   if (err != nil) {
  //     fmt.Println(err)
  //   } else {
  //     fmt.Println(hasil)
  //   }

  //   db, err := sql.Open("mysql", database.ConnectDatabase(username, password, host, port, databasetype))
  //   if (err != nil) {
  //     fmt.Printf("Error %s while opening database\n", err)
  //   }

  //   pingErr := db.Ping()
  //   if pingErr != nil {
  //       log.Fatal(pingErr)
  //   }
  //   fmt.Println("Connected!")
  //   defer db.Close()

