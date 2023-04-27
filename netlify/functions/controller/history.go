package controller

import (
	"TUBES3_13521102/netlify/functions/repository"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func GetHistoryByHistoryID(c *gin.Context){
  // Receive historyID through front end
  historyID, err := strconv.ParseInt(c.Param("historyID"), 10, 64)
  if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
          "error": "Invalid history ID",
      })
      return
  }

  db, err := sql.Open("mysql", "root:PNGO6atNekbjjq4g2yPy@tcp(containers-us-west-13.railway.app:6330)/railway")
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

  historybyID, err := repository.GetHistoryByHistoryID(db, historyID)
  // historyByID, err := repository.GetHistoryByHistoryID(database.db, historyID)
  c.JSON(http.StatusOK, gin.H{
      "historybyID": historybyID,
  })
}


// func InsertHistory(c *gin.Context){
//   var (
//     history structs.History
//     err error
//   )
//   db, err := sql.Open("mysql", "root:PNGO6atNekbjjq4g2yPy@tcp(containers-us-west-13.railway.app:6330)/railway")
//   if err != nil {
//       c.JSON(http.StatusInternalServerError, gin.H{
//           "error": err.Error(),
//       })
//       return
//   }
//   pingErr := db.Ping()
//   if pingErr != nil {
//       log.Fatal(pingErr)
//   }

//   defer db.Close()

//   err = c.ShouldBindJSON(&history)
//   if (err != nil){
//     panic(err)
//   }
//   err = repository.InsertHistory(db, history, history.)
// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"result": "Success Create History",
// 	})

// }
