package main

import (
	"TUBES3_13521102/netlify/functions/controller"
	"TUBES3_13521102/netlify/functions/database"
	"database/sql"
	"fmt"
	"log"
  "net/http"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

// Variables to connect to database host
const (
  username = "root"
  host = "containers-us-west-13.railway.app"
  password = "PNGO6atNekbjjq4g2yPy"
  port = "6330"
  databasetype = "railway"
)


var (
  db *sql.DB
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	db, err := sql.Open("mysql", database.ConnectDatabase(username, password, host, port, databasetype))
	if err != nil {
		fmt.Printf("Error %s while opening database\n", err)
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       pingErr.Error(),
		}, nil
	}
	fmt.Println("Connected!")
	defer db.Close()
	
	switch request.HTTPMethod {
	case http.MethodGet:
		if request.Resource == "/history" {
      response, err := controller.GetAllHistoryMessage(request)
      if err != nil {
        return response, err
      }
      response.Headers["Access-Control-Allow-Origin"] = "*"
      return response, nil
    }
	case http.MethodPost:
		if request.Resource == "/getmessage" {
			return controller.ParseUserMessage(request)
		}
	case http.MethodOptions:
		if request.Resource == "/getmessage" {
			headers := map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST",
				"Access-Control-Allow-Headers": "Content-Type",
			}
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Headers:    headers,
			}, nil
		}
	default:
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       http.StatusText(http.StatusMethodNotAllowed),
		}, nil
	}

    return &events.APIGatewayProxyResponse{
      StatusCode: http.StatusNotImplemented,
      Body:       http.StatusText(http.StatusNotImplemented),
    }, nil
}

func main() {
  // Make the handler available for Remote Procedure Call
  lambda.Start(handler)
}
