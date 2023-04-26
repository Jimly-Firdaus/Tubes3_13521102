package controller

import (
	"TUBES3_13521102/src-backend/repository"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func getAllUserMessageHandler(r *gin.Engine){
  r.GET("/userMessage", func(c *gin.Context) {
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

    userMessageList, err := repository.GetAllUserMessage(db)
    c.JSON(http.StatusOK, gin.H{
        "userMessage": userMessageList,
    })
  })
}

func getUserMessageByIDHandle(r *gin.Engine){
  r.GET("/userMessage/:userQuestionID", func(c *gin.Context) {
    // Receive historyID through front end
    userQuestionID, err := strconv.ParseInt(c.Param("userQuestionID"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid user question ID",
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

    userMessagebyID, err := repository.GetHistoryMessageByID(db, userQuestionID)
    c.JSON(http.StatusOK, gin.H{
        "userMessagebyID": userMessagebyID,
    })
  })

}
