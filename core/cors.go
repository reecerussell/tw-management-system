package core

import (
	"github.com/aws/aws-lambda-go/events"
)

// CORSHeaders returns a map of headers used for CORS.
func CORSHeaders(req events.APIGatewayProxyRequest) map[string]string {
	var origin string

	if v, ok := req.StageVariables["CORS_ORIGIN"]; ok {
		origin = v
	} else {
		origin = "*"
	}

	return map[string]string{
		"Access-Control-Allow-Origin":  origin,
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		"Access-Control-Allow-Methods": req.HTTPMethod,
	}
}
