package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	department := req.PathParameters["department"]
	if department == "" {
		resp := core.NewErrorWithStatus(
			fmt.Errorf("missing department"),
			http.StatusBadRequest,
		).Response()
		resp.Headers = core.CORSHeaders(http.MethodGet)
		return resp, nil
	}

	log.Printf("Getting queue buster for department '%s'", department)

	qb, err := queueBusters.Get(department)
	if err != nil {
		resp := err.Response()
		resp.Headers = core.CORSHeaders(http.MethodGet)
		return resp, nil
	}

	data, _ := json.Marshal(qb)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    core.CORSHeaders(http.MethodGet),
		Body:       string(data),
	}, nil
}

func main() {
	lambda.Start(Handle)
}
