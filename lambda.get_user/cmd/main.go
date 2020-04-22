package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/reecerussell/tw-management-system/core/users"
)

var repo users.Repository

func init() {
	repo = users.NewRepository()
}

func GetUserHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.HTTPMethod != http.MethodGet {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       http.StatusText(http.StatusMethodNotAllowed),
		}, nil
	}

	id := req.QueryStringParameters["id"]
	if id == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       http.StatusText(http.StatusNotFound),
		}, nil
	}

	u, err := repo.Get(id)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	bytes, _ := json.Marshal(u)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bytes),
	}, nil
}

func main() {
	lambda.Start(GetUserHandler)
}
