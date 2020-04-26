package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	usrs "github.com/reecerussell/tw-management-system/core/users"
	"github.com/reecerussell/tw-management-system/core/users/usecase"
)

var users usecase.UserUsecase

func init() {
	users = usrs.Usecase()
}

func HandleGetUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Path: %v\n", req.Path)
	log.Printf("Path Parameters: %d\n", len(req.PathParameters))
	for k, v := range req.PathParameters {
		log.Printf("%s: %s\n", k, v)
	}
	log.Printf("Query Strings: %d\n", len(req.QueryStringParameters))
	for k, v := range req.QueryStringParameters {
		log.Printf("%s: %s\n", k, v)
	}
	log.Printf("Method: %s\n", req.HTTPMethod)

	id := req.PathParameters["id"]
	if id == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       http.StatusText(http.StatusNotFound),
		}, nil
	}

	user, err := users.Get(id)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: err.Status(),
			Body:       err.Message(),
		}, nil
	}

	json, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func main() {
	lambda.Start(HandleGetUser)
}
