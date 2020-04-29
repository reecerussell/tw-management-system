# Create Queue Buster

This lambda function is used to create new queue buster records. Triggered by an HTTP POST request, via an API Gateway.

The request requires JSON data in the request in the format of:

```json
{
	"department": "<department name>",
	"status": true
}
```

A successful response will return an OK (200) status, with an empty body.

**NOTE:** this lambda function should require authorization via the API Gateway.

## Permissions

This lambda function requires full read/write access to DynamoDB.
