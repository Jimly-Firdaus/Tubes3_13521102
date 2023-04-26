package controller

import (
	"TUBES3_13521102/src-backend/repository"
	"TUBES3_13521102/src-backend/structs"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func InsertUserMessage(c *gin.Context){
  var (
    msg structs.Message
    err error
  )

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


  err = c.ShouldBindJSON(&msg)
  if (err != nil){
    panic(err)
  }
  err = repository.InsertUserMessage(db, msg)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Create History",
	})

}

func GetUserMessageByID(c *gin.Context){
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

   userMessage, err := repository.GetUserMessageByID(db, userQuestionID)
   // historyByID, err := repository.GetHistoryByHistoryID(database.db, historyID)
   c.JSON(http.StatusOK, gin.H{
       "userMessage": userMessage,
   })
}
func ParseUserMessage(c* gin.Context) {
  // Parse the request body into a Request struct
  var req structs.Request
  if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

  // Use the message and method fields in the Request struct
  fmt.Println(req.Message)
  fmt.Println(req.Method)

  // Write a response back to the client
  // ...
  resp := gin.H{"message": "Received request successfully"}
	c.JSON(http.StatusOK, resp)
}
// func GetAllUserMessageHandler(r *gin.Engine){
//   r.GET("/userMessage", func(c *gin.Context) {
//     db, err := sql.Open("mysql", "root:PNGO6atNekbjjq4g2yPy@tcp(containers-us-west-13.railway.app:6330)/railway")
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{
//             "error": err.Error(),
//         })
//         return
//     }
//     pingErr := db.Ping()
//     if pingErr != nil {
//         log.Fatal(pingErr)
//     }

//     defer db.Close()

//   //   userMessageList, err := repository.GetAllUserMessage(db)
//   //   c.JSON(http.StatusOK, gin.H{
//   //       "userMessage": userMessageList,
//   //   })
//   // })
// }

// func getUserMessageByIDHandle(r *gin.Engine){
//   r.GET("/userMessage/:userQuestionID", func(c *gin.Context) {
//     // Receive historyID through front end
//     userQuestionID, err := strconv.ParseInt(c.Param("userQuestionID"), 10, 64)
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{
//             "error": "Invalid user question ID",
//         })
//         return
//     }

//     db, err := sql.Open("mysql", "root:PNGO6atNekbjjq4g2yPy@tcp(containers-us-west-13.railway.app:6330)/railway")
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{
//             "error": err.Error(),
//         })
//         return
//     }
//     pingErr := db.Ping()
//     if pingErr != nil {
//         log.Fatal(pingErr)
//     }

//     defer db.Close()

//     userMessagebyID, err := repository.GetHistoryMessageByID(db, userQuestionID)
//     c.JSON(http.StatusOK, gin.H{
//         "userMessagebyID": userMessagebyID,
//     })
//   })

// }
