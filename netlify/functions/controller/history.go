package controller

import (
	"TUBES3_13521102/netlify/functions/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func GetHistoryByID(request events.APIGatewayProxyRequest, db *sql.DB) (*events.APIGatewayProxyResponse, error) {
	// Parse the request body into a Request struct
	type Hid struct {
		Id int64
	}
	var hid Hid
	var status string
	if err := json.Unmarshal([]byte(request.Body), &hid); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	// Use the message and method fields in the Request struct
	fmt.Println(hid.Id)
	status = "200"
	hist, err := repository.GetHistoryByHistoryID(db, hid.Id)
	if err != nil {
		panic(err)
	}
	// Parameter jadi struct.
	// TO DO : Masukin database
	// Write a response back to the client
	responseBody, err := json.Marshal(map[string]interface{}{
		"message": status,
		"history": hist,
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
