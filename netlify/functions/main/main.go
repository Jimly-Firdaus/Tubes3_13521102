package main

import (
	"TUBES3_13521102/netlify/functions/main/controller"
	"TUBES3_13521102/netlify/functions/main/database"
	"database/sql"
	"fmt"
	"log"
  "net/http"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/akrylysov/algnhsa"
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
    // Create http.ResponseWriter and *http.Request from events.APIGatewayProxyRequest
    w, r, err := algnhsa.NewResponseWriter(request), algnhsa.NewRequest(request)
    if err != nil {
        return &events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       err.Error(),
        }, nil
    }

    // Your server-side functionality
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

    switch r.Method {
    case http.MethodGet:
        if r.URL.Path == "/history" {
            // Handle GET /history request
            controller.GetAllHistoryMessage(w, r)
        }
    case http.MethodPost:
        if r.URL.Path == "/message" {
            // Handle POST /message request
            // ...
        } else if r.URL.Path == "/getmessage" {
            // Handle POST /getmessage request
            controller.ParseUserMessage(w, r)
        }
    case http.MethodOptions:
        if r.URL.Path == "/getmessage" {
            // Handle OPTIONS /getmessage request
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "POST")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            w.WriteHeader(http.StatusOK)
            return algnhsa.ProxyResponse(w), nil
        }
    }

    return algnhsa.ProxyResponse(w), nil
}

func main() {
  // Make the handler available for Remote Procedure Call
  lambda.Start(handler)
}
