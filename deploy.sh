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
env GOOS=linux GOARCH=amd64 go build -o /tmp/$AUTHORIZER_NAME lambda.authorizer/main.go
zip -j /tmp/$AUTHORIZER_NAME.zip /tmp/$AUTHORIZER_NAME
aws lambda update-function-code --function-name $AUTHORIZER_NAME --zip-file fileb:///tmp/$AUTHORIZER_NAME.zip

# Login
env GOOS=linux GOARCH=amd64 go build -o /tmp/$LOGIN_NAME lambda.login/main.go
zip -j /tmp/$LOGIN_NAME.zip /tmp/$LOGIN_NAME
aws lambda update-function-code --function-name $LOGIN_NAME --zip-file fileb:///tmp/$LOGIN_NAME.zip

# Change Password
env GOOS=linux GOARCH=amd64 go build -o /tmp/$CHANGE_PASSWORD_NAME lambda.change-password/main.go
zip -j /tmp/$CHANGE_PASSWORD_NAME.zip /tmp/$CHANGE_PASSWORD_NAME
aws lambda update-function-code --function-name $CHANGE_PASSWORD_NAME --zip-file fileb:///tmp/$CHANGE_PASSWORD_NAME.zip

# Create User
env GOOS=linux GOARCH=amd64 go build -o /tmp/$CREATE_USER_NAME lambda.create-user/main.go
zip -j /tmp/$CREATE_USER_NAME.zip /tmp/$CREATE_USER_NAME
aws lambda update-function-code --function-name $CREATE_USER_NAME --zip-file fileb:///tmp/$CREATE_USER_NAME.zip

# Delete User
env GOOS=linux GOARCH=amd64 go build -o /tmp/$DELETE_USER_NAME lambda.delete-user/main.go
zip -j /tmp/$DELETE_USER_NAME.zip /tmp/$DELETE_USER_NAME
aws lambda update-function-code --function-name $DELETE_USER_NAME --zip-file fileb:///tmp/$DELETE_USER_NAME.zip

# Get User
env GOOS=linux GOARCH=amd64 go build -o /tmp/$GET_USER_NAME lambda.get-user/main.go
zip -j /tmp/$GET_USER_NAME.zip /tmp/$GET_USER_NAME
aws lambda update-function-code --function-name $GET_USER_NAME --zip-file fileb:///tmp/$GET_USER_NAME.zip

# Get Users
env GOOS=linux GOARCH=amd64 go build -o /tmp/$GET_USERS_NAME lambda.get-users/main.go
zip -j /tmp/$GET_USERS_NAME.zip /tmp/$GET_USERS_NAME
aws lambda update-function-code --function-name $GET_USERS_NAME --zip-file fileb:///tmp/$GET_USERS_NAME.zip

# Update User
env GOOS=linux GOARCH=amd64 go build -o /tmp/$UPDATE_NAME lambda.update-user/main.go
zip -j /tmp/$UPDATE_NAME.zip /tmp/$UPDATE_NAME
aws lambda update-function-code --function-name $UPDATE_NAME --zip-file fileb:///tmp/$UPDATE_NAME.zip

# Create Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$CREATE_QB_NAME lambda.create-queue-buster/main.go
zip -j /tmp/$CREATE_QB_NAME.zip /tmp/$CREATE_QB_NAME
aws lambda update-function-code --function-name $CREATE_QB_NAME --zip-file fileb:///tmp/$CREATE_QB_NAME.zip

# Delete Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$DELETE_QB_NAME lambda.delete-queue-buster/main.go
zip -j /tmp/$DELETE_QB_NAME.zip /tmp/$DELETE_QB_NAME
aws lambda update-function-code --function-name $DELETE_QB_NAME --zip-file fileb:///tmp/$DELETE_QB_NAME.zip

# Get Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$GET_QB_NAME lambda.get-queue-buster/main.go
zip -j /tmp/$GET_QB_NAME.zip /tmp/$GET_QB_NAME
aws lambda update-function-code --function-name $GET_QB_NAME --zip-file fileb:///tmp/$GET_QB_NAME.zip

# Get Queue Busters
env GOOS=linux GOARCH=amd64 go build -o /tmp/$GET_QBS_NAME lambda.get-queue-busters/main.go
zip -j /tmp/$GET_QBS_NAME.zip /tmp/$GET_QBS_NAME
aws lambda update-function-code --function-name $GET_QBS_NAME --zip-file fileb:///tmp/$GET_QBS_NAME.zip

# Enable Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$ENABLE_QB_NAME lambda.enable-queue-buster/main.go
zip -j /tmp/$ENABLE_QB_NAME.zip /tmp/$ENABLE_QB_NAME
aws lambda update-function-code --function-name $ENABLE_QB_NAME --zip-file fileb:///tmp/$ENABLE_QB_NAME.zip

# Disable Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$DISABLE_QB_NAME lambda.disable-queue-buster/main.go
zip -j /tmp/$DISABLE_QB_NAME.zip /tmp/$DISABLE_QB_NAME
aws lambda update-function-code --function-name $DISABLE_QB_NAME --zip-file fileb:///tmp/$DISABLE_QB_NAME.zip