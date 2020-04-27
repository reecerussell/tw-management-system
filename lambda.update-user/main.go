package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

func HandleUpdate(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user dto.User

	body := make([]byte, base64.StdEncoding.DecodedLen(len(req.Body)))
	log.Printf("Body: %v\n", req.Body)
	if req.IsBase64Encoded {
		base64.StdEncoding.Decode(body, []byte(req.Body))
	} else {
		body = []byte(req.Body)
	}

	buf := bytes.NewBuffer(body)
	_ = json.NewDecoder(buf).Decode(&user)

	if user.ID == "" {
		err := core.NewErrorWithStatus(fmt.Errorf("missing 'id' value"), http.StatusBadRequest)
		resp := err.Response()
		resp.Headers = core.CORSHeaders(http.MethodPut)
		return resp, nil
	}

	err := users.Update(&user)
	if err != nil {
		resp := err.Response()
		resp.Headers = core.CORSHeaders(http.MethodPut)
		return resp, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    core.CORSHeaders(http.MethodPut),
	}, nil
}

func main() {
	lambda.Start(HandleUpdate)
}
