package main

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/jwt"
)

var (
	publicKey *rsa.PublicKey
	errLog    *log.Logger
)

func init() {
	s, err := core.NewSecret(os.Getenv("SECRET_NAME"))
	if err != nil {
		panic(err)
	}

	pk, err := s.RSAPublicKey("public")
	if err != nil {
		panic(err)
	}

	publicKey = pk
	errLog = log.New(os.Stderr, "[AUTHORIZER][ERROR]: ", log.LstdFlags)
}

// HandleAuthorize is a Handler function for lambda.
func HandleAuthorize(evt events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	parts := strings.Split(evt.AuthorizationToken, " ")
	if len(parts) < 2 {
		errLog.Printf("invalid token format: %s", evt.AuthorizationToken)
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	if parts[0] != "Bearer" {
		errLog.Printf("invalid token scheme: %s", parts[0])
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	token, err := jwt.FromToken([]byte(parts[1]))
	if err != nil {
		errLog.Printf("token error: %v", err)
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	valid, err := token.Check(publicKey)
	if err != nil {
		errLog.Printf("invalid token: %v", err)
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	if !valid {
		errLog.Printf("token not valid")
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	log.Printf("[AUTHORIZER]: success: valid token")

	return generatePolicy("user", "Allow"), nil
}

func generatePolicy(principalID, effect string) events.APIGatewayCustomAuthorizerResponse {
	res := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}

	if effect != "" {
		res.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action: []string{"execute-api:Invoke"},
					Effect: effect,
					Resource: []string{
						fmt.Sprintf("arn:aws:execute-api:%s:%s:%s/*/*/*",
							os.Getenv("AWS_REGION"),
							os.Getenv("ACCOUNT_ID"),
							os.Getenv("API_ID")),
					},
				},
			},
		}
	}

	return res
}

func main() {
	lambda.Start(HandleAuthorize)
}
