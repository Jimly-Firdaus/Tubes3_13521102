package controller

import (
	"TUBES3_13521102/src-backend/repository"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
func GetAllHistoryMessage(c *gin.Context){
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

  historyMessageList, err := repository.GetAllHistoryMessage(db)
  c.JSON(http.StatusOK, gin.H{
      "historyMessage": historyMessageList,
  })
}
