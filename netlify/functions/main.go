package main

import (
	"TUBES3_13521102/netlify/functions/controller"
	"TUBES3_13521102/netlify/functions/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	// "github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

// Variables to connect to database host
const (
	username     = "root"
	host         = "containers-us-west-13.railway.app"
	password     = "PNGO6atNekbjjq4g2yPy"
	port         = "6330"
	databasetype = "railway"
)

var (
	db      *sql.DB
	regexes []*regexp.Regexp
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	patterns := []string{
		`^(?i)Tambahkan pertanyaan (.*) dengan jawaban (.*)$`, // Tambah pertanyaan
		`^(?i)Hapus pertanyaan (.*)$`,                         // Hapus pertanyaan
		`(?i)^(Hari apa )?[0-9]{2}/[0-9]{2}/[0-9]{4}\??$`,     // Kalendar
		`^[\d()+\-*\/.^ ]+$`,                                  // Kalkulator
		`.*`,                                                  // Pertanyaan Teks
	}
	// Compile the patterns into regex objects
	regexes := make([]*regexp.Regexp, len(patterns))
	for i, pattern := range patterns {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		regexes[i] = regex
	}

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
		fmt.Printf("Request: %+v\n", request)
		fmt.Println("Hit history get")
		if request.Path == "/.netlify/functions/endpoint/history" {
			response, err := controller.GetAllHistoryMessage(request, db)
			if err != nil {
				return response, err
			}
			response.Headers["Access-Control-Allow-Origin"] = "*"
			fmt.Println("Response headers: %v", response.Headers)
			fmt.Println("Response body: %s", response.Body)
			fmt.Println("Hit history")
			return response, nil
		}
		//  else if request.Path == ".netlify/functions/endpoint/giveresponse" {
		// 	response, err := controller.GetAllHistoryMessage(request)
		// }
	case http.MethodPost:
		if request.Path == "/.netlify/functions/endpoint/getmessage" {
			return controller.ParseUserMessage(request, db, regexes)
		}
	case http.MethodOptions:
		if request.Path == "/.netlify/functions/endpoint/getmessage" {
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
		fmt.Println("HTTP method: %s", request.HTTPMethod)
		fmt.Println("Resource: %s", request.Resource)
		fmt.Println("Hit default")
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       http.StatusText(http.StatusMethodNotAllowed),
		}, nil
	}
	fmt.Println("HTTP method: %s", request.HTTPMethod)
	fmt.Println("Resource: %s", request.Resource)
	fmt.Println("Hit outside switch")
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotImplemented,
		Body:       http.StatusText(http.StatusNotImplemented),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)

	// expression := "1/8"

	// fmt.Println(FeatureCalculator.CalculateExpression(expression))
}
