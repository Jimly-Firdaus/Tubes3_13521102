package controller

import (
	"TUBES3_13521102/netlify/functions/repository"
	"TUBES3_13521102/netlify/functions/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gin-gonic/gin"
)

func InsertUserMessage(c *gin.Context) {
	var (
		msg structs.Request
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
	if err != nil {
		panic(err)
	}
	err = repository.InsertMessages(db, msg)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Create History",
	})

}

func GetUserMessageByID(c *gin.Context) {
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

	userMessage, err := repository.GetMessagesByID(db, userQuestionID)
	if err != nil {
		panic(err)
	}
	// historyByID, err := repository.GetHistoryByHistoryID(database.db, historyID)
	c.JSON(http.StatusOK, gin.H{
		"userMessage": userMessage,
	})
}
func ParseUserMessage(request events.APIGatewayProxyRequest, db *sql.DB, regex []*regexp.Regexp) (*events.APIGatewayProxyResponse, error) {
	// Parse the request body into a Request struct
	var req structs.Request
	var status string
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	// Use the message and method fields in the Request struct
	fmt.Println(req.Id)
	fmt.Println(req.Text)
	fmt.Println(req.Response)
	fmt.Println(req.SentTime)
	fmt.Println(req.HistoryId)
	fmt.Println(req.HistoryTimeStamp)
	fmt.Println(req.Method)
	status = "200"
	FilterMessage(&req, &status, db, regex)
	// Parameter jadi struct.
	// TO DO : Masukin database
	// Write a response back to the client
	responseBody, err := json.Marshal(map[string]interface{}{
		"message":     status,
		"botResponse": req,
	})
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

