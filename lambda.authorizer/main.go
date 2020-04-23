package main

import (
	"crypto/rsa"
	"errors"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/jwt"
)

var publicKey *rsa.PublicKey

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
}

// HandleAuthorize is a Handler function for lambda.
func HandleAuthorize(evt events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	parts := strings.Split(evt.AuthorizationToken, " ")
	if len(parts) < 2 {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	if parts[0] != "Bearer" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	token, err := jwt.FromToken([]byte(parts[1]))
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	valid, err := token.Check(publicKey)
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	if !valid {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	return generatePolicy("user", "Allow", evt.MethodArn), nil
}

func generatePolicy(principalID, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	res := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}

	if effect != "" && resource != "" {
		res.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	return res
}

func main() {
	lambda.Start(HandleAuthorize)
}
