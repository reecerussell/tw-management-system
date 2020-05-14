package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/reecerussell/tw-management-system/core"
	usrs "github.com/reecerussell/tw-management-system/core/users"
	"github.com/reecerussell/tw-management-system/core/users/dto"
	"github.com/reecerussell/tw-management-system/core/users/usecase"
)

var users usecase.UserUsecase

func init() {
	users = usrs.Usecase()
}

func HandleChangePassword(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data dto.ChangePassword

	body := make([]byte, base64.StdEncoding.DecodedLen(len(req.Body)))
	log.Printf("Body: %v\n", req.Body)
	if req.IsBase64Encoded {
		base64.StdEncoding.Decode(body, []byte(req.Body))
	} else {
		body = []byte(req.Body)
	}

	buf := bytes.NewBuffer(body)
	_ = json.NewDecoder(buf).Decode(&data)

	err := users.ChangePassword(&data)
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
	lambda.Start(HandleChangePassword)
}
