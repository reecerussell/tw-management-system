package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/queuebuster"
	"github.com/reecerussell/tw-management-system/core/queuebuster/dto"
	"github.com/reecerussell/tw-management-system/core/queuebuster/usecase"
)

var queueBusters usecase.QueueBusterUsecase

func init() {
	queueBusters = queuebuster.Usecase()
}

// Handle handles HTTP POSt requests to create a queue buster record.
func Handle(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data dto.QueueBuster

	body := make([]byte, base64.StdEncoding.DecodedLen(len(req.Body)))
	if req.IsBase64Encoded {
		base64.StdEncoding.Decode(body, []byte(req.Body))
	} else {
		body = []byte(req.Body)
	}

	buf := bytes.NewBuffer(body)
	_ = json.NewDecoder(buf).Decode(&data)

	err := queueBusters.Create(&data)
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
