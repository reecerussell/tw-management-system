package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-lambda-go/events"
	usrs "github.com/reecerussell/tw-management-system/core/users"
	"github.com/reecerussell/tw-management-system/core/users/usecase"
)

var users usecase.UserUsecase

func init() {
	users = usrs.Usecase()
}

func HandleGetUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := req.PathParameters["id"]
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
			Body:       http.StatusText(http.StatusNotFound),
		}, nil
	}

	json, _ := json.Marshal(user)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func Handle(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch true {
	case req.Path == "/" && req.HTTPMethod == http.MethodGet:
		return HandleGetUser(req)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadGateway,
		Body:       http.StatusText(http.StatusBadGateway),
	}, nil
}

func main() {
	lambda.Start(Handle)
}
