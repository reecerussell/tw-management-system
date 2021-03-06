package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/reecerussell/tw-management-system/core"
	usrs "github.com/reecerussell/tw-management-system/core/users"
	"github.com/reecerussell/tw-management-system/core/users/usecase"
)

var users usecase.UserUsecase

func init() {
	users = usrs.Usecase()
}

func HandleGetUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	users, err := users.All()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: err.Status(),
			Headers:    core.CORSHeaders(req),
			Body:       http.StatusText(http.StatusNotFound),
		}, nil
	}

	json, _ := json.Marshal(users)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    core.CORSHeaders(req),
		Body:       string(json),
	}, nil
}

func main() {
	lambda.Start(HandleGetUser)
}
