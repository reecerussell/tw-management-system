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
	id := req.PathParameters["id"]
	if id == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Headers:    core.CORSHeaders(req),
			Body:       http.StatusText(http.StatusNotFound),
		}, nil
	}

	user, err := users.Get(id)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: err.Status(),
			Headers:    core.CORSHeaders(req),
			Body:       err.Message(),
		}, nil
	}

	json, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    core.CORSHeaders(req),
		Body:       string(json),
	}, nil
}

func main() {
	lambda.Start(HandleGetUser)
}
