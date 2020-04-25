# Login

Login is a simple lambda function which is used to generate an OAuth access token for a user. This requires no authentication.

## Variables

-   `AUTH_SECRET_NAME`: this variable is used to create the secret for the token signing.
-   `JWT_ISSUER`: sets the access token issuer claim.
-   `JWT_AUDIENCE`: sets the access token audience claim.

## Permissions

This function requires read writes to the secret manager.
