package controller

import (
	"TUBES3_13521102/netlify/functions/main/repository"
	"TUBES3_13521102/netlify/functions/main/structs"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
  "encoding/json"

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
func ParseUserMessage(w http.ResponseWriter, r *http.Request) {
  // Parse the request body into a Request struct
  var req structs.Request
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
  }

  // Use the message and method fields in the Request struct
  // fmt.Println(req.Message)
  fmt.Println(req.Id)
  fmt.Println(req.Text)
  fmt.Println(req.Response)
  fmt.Println(req.SentTime)
  fmt.Println(req.HistoryId)
  fmt.Println(req.HistoryTimeStamp)
  fmt.Println(req.Method)

  // Write a response back to the client
  // ...
  resp := map[string]interface{}{"message": "Received request successfully"}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(resp)
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
