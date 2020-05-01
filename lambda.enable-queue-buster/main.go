package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/queuebuster"
	"github.com/reecerussell/tw-management-system/core/queuebuster/usecase"
)

var queueBusters usecase.QueueBusterUsecase

func init() {
	queueBusters = queuebuster.Usecase()
}

// Handle is the lambda function handler.
func Handle(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	department, _ := url.QueryUnescape(req.PathParameters["department"])
	if department == "" {
		resp := core.NewErrorWithStatus(
			fmt.Errorf("missing department"),
			http.StatusBadRequest,
		).Response()
		resp.Headers = core.CORSHeaders(http.MethodPost)
		return resp, nil
	}

	err := queueBusters.Enable(department)
	if err != nil {
		resp := err.Response()
		resp.Headers = core.CORSHeaders(http.MethodPost)
		return resp, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    core.CORSHeaders(http.MethodPost),
	}, nil
}

func main() {
	lambda.Start(Handle)
}
