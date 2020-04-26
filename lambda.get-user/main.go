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
	log.Println("Path Parameters")
	for k, v := range req.PathParameters {
		log.Printf("%s: %s", k, v)
	}

	log.Println("Query Parameters")
	for k, v := range req.QueryStringParameters {
		log.Printf("%s: %s", k, v)
	}

	id, ok := req.PathParameters["id"]
	log.Printf("ID: %s, %v\n", id, ok)

	if !ok || id == "" {
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
