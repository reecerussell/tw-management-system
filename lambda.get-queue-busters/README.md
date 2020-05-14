# Get Queue Busters

A lambda function used to handle HTTP GET requests to reteive all queue busters from the database.

A successful response will have an OK (200) status code, with an array of objects in the body, in JSON format. An example of this is shown below:

```json
[
	{
		"department": "Example Department 1",
		"enabled": true
	},
	{
		"department": "Example Department 2",
		"enabled": false
	}
]
```

**NOTE:** this lambda function should require authorization via the API Gateway.

## Permissions

This lambda function requires read access to DynamoDB.
