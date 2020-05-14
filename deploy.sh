#!/bin/bash

#
# Builds and deploys all of the lambda functions and assumes they all exist.
#

AUTHORIZER_NAME=tw-authorizer
LOGIN_NAME=tw-login
CHANGE_PASSWORD_NAME=tw-change-password
CREATE_USER_NAME=tw-create-user
DELETE_USER_NAME=tw-delete-user
GET_USER_NAME=tw-get-user
GET_USERS_NAME=tw-get-users
UPDATE_NAME=tw-update-user
CREATE_QB_NAME=tw-create-queue-buster
DELETE_QB_NAME=tw-delete-queue-buster
GET_QB_NAME=tw-get-queue-buster
GET_QBS_NAME=tw-get-queue-busters
ENABLE_QB_NAME=tw-enable-queue-buster
DISABLE_QB_NAME=tw-disable-queue-buster

# Authorizer
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.authorizer/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $AUTHORIZER_NAME --zip-file fileb:///tmp/main.zip

# Login
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.login/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $LOGIN_NAME --zip-file fileb:///tmp/main.zip

# Change Password
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.change-password/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $CHANGE_PASSWORD_NAME --zip-file fileb:///tmp/main.zip

# Create User
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.create-user/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $CREATE_USER_NAME --zip-file fileb:///tmp/main.zip

# Delete User
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.delete-user/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $DELETE_USER_NAME --zip-file fileb:///tmp/main.zip

# Get User
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.get-user/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $GET_USER_NAME --zip-file fileb:///tmp/main.zip

# Get Users
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.get-users/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $GET_USERS_NAME --zip-file fileb:///tmp/main.zip

# Update User
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.update-user/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $UPDATE_NAME --zip-file fileb:///tmp/main.zip

# Create Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.create-queue-buster/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $CREATE_QB_NAME --zip-file fileb:///tmp/main.zip

# Delete Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.delete-queue-buster/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $DELETE_QB_NAME --zip-file fileb:///tmp/main.zip

# Get Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.get-queue-buster/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $GET_QB_NAME --zip-file fileb:///tmp/main.zip

# Get Queue Busters
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.get-queue-busters/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $GET_QBS_NAME --zip-file fileb:///tmp/main.zip

# Enable Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.enable-queue-buster/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $ENABLE_QB_NAME --zip-file fileb:///tmp/main.zip

# Disable Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/main lambda.disable-queue-buster/main.go
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name $DISABLE_QB_NAME --zip-file fileb:///tmp/main.zip