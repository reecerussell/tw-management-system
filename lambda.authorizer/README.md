# Authorizer

Authorizer is used as the Lambda Auth Func to authorize requests to protected resources.

## Variables

-   `SECRET_NAME` is an environment variable used to specify the secret used to verify the JWT tokens. **NOTE:** this secret must have a valid RSA public key in PEM format, stored in the "public" field in the secret/
