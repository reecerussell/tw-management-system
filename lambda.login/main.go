package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"

	usrs "github.com/reecerussell/tw-management-system/core/users"
	"github.com/reecerussell/tw-management-system/core/users/dto"
	"github.com/reecerussell/tw-management-system/core/users/usecase"

	"github.com/aws/aws-lambda-go/events"
)

var users usecase.UserUsecase

func init() {
	users = usrs.Usecase()
}

func HandleLogin(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.HTTPMethod != http.MethodPost {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       http.StatusText(http.StatusMethodNotAllowed),
		}, nil
	}

	var creds dto.UserCredentials
	rdr := strings.NewReader(req.Body)
	_ = json.NewDecoder(rdr).Decode(&creds)

	ac, err := users.Login(&creds)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: err.Status(),
			Body:       err.Message(),
		}, nil
	}

	data, _ := json.Marshal(ac)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(data),
	}, nil
}

func main() {
	lambda.Start(HandleLogin)
}
