package core

import (
	"os"
	"strings"
)

// CORSOrigin is the allowed origin value in the cors headers.
var CORSOrigin string

func init() {
	if v := os.Getenv("CORS_ORIGIN"); v != "" {
		CORSOrigin = v
	} else {
		CORSOrigin = "*"
	}
}

// CORSHeaders returns a map of headers used for CORS.
func CORSHeaders(methods ...string) map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":  CORSOrigin,
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		"Access-Control-Allow-Methods": strings.Join(methods, ","),
	}
}
