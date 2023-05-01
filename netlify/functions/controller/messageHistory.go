package controller

import (
	"TUBES3_13521102/netlify/functions/repository"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func GetAllHistoryMessage(request events.APIGatewayProxyRequest, db *sql.DB) (*events.APIGatewayProxyResponse, error) {
	// db, err := sql.Open("mysql", "root:PNGO6atNekbjjq4g2yPy@tcp(containers-us-west-13.railway.app:6330)/railway")
	// if err != nil {
	// 	return &events.APIGatewayProxyResponse{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Body:       err.Error(),
	// 	}, nil
	// }
	// defer db.Close()

	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// 	return &events.APIGatewayProxyResponse{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Body:       pingErr.Error(),
	// 	}, nil
	// }

	historyMessageList, err := repository.GetAllHistoryMessage(db)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	responseBody, err := json.Marshal(map[string]interface{}{
		"historyMessage": historyMessageList,
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
