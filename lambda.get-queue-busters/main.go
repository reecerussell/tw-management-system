package main

import (
	"encoding/json"
	"net/http"

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
	qbs, err := queueBusters.GetAll()
	if err != nil {
		resp := err.Response()
		resp.Headers = core.CORSHeaders(req)
		return resp, nil
	}

	data, _ := json.Marshal(qbs)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    core.CORSHeaders(req),
		Body:       string(data),
	}, nil
}

func main() {
	lambda.Start(Handle)
}
