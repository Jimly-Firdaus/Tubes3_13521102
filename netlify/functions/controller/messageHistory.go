package controller

import (
	"TUBES3_13521102/netlify/functions/repository"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func GetAllHistoryMessage(request events.APIGatewayProxyRequest, db *sql.DB) (*events.APIGatewayProxyResponse, error) {
	historyPayload, err := repository.GetAllHistory(db)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	responseBody, err := json.Marshal(map[string]interface{}{
		"historyPayload": historyPayload,
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
