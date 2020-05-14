package main

import (
	"fmt"
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

func HandleUpdate(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]
	if id == "" {
		err := core.NewErrorWithStatus(fmt.Errorf("missing 'id' value"), http.StatusBadRequest)
		resp := err.Response()
		resp.Headers = core.CORSHeaders(req)
		return resp, nil
	}

	err := users.Delete(id)
	if err != nil {
		resp := err.Response()
		resp.Headers = core.CORSHeaders(req)
		return resp, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    core.CORSHeaders(req),
	}, nil
}

func main() {
	lambda.Start(HandleUpdate)
}
