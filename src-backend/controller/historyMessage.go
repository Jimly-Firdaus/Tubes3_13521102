package controller

import (
	"TUBES3_13521102/src-backend/repository"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func getAllHistoryMessageHandler(r *gin.Engine){
  r.GET("/historyMessage", func(c *gin.Context) {
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
  })
}

func getHistoryMessageByHistoryIDHandler(r *gin.Engine){
  r.GET("/historyMessage/:historyID", func(c *gin.Context) {
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

    historyMessagebyID, err := repository.GetHistoryMessageByID(db, historyID)
    c.JSON(http.StatusOK, gin.H{
        "historyMessagebyID": historyMessagebyID,
    })
  })
}
