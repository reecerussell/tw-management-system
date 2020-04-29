# Get Queue Busters

A lambda function used to handle HTTP GET requests to reteive a specific queue busters from the database, specified by a department.

A successful response will have an OK (200) status code, with the body populated with JSON data of the requested queue buster.

```json
{
	"department": "Example Department 1",
	"enabled": true
}
```

**NOTE:** this lambda function should require authorization via the API Gateway.

## Permissions

This lambda function requires read access to DynamoDB.
