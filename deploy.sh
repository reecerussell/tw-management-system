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
ENABLE_ANNOUNCEMENTS_QB_NAME=tw-enable-queue-buster-announcements
DISABLE_ANNOUNCEMENTS_QB_NAME=tw-disbale-queue-buster-announcements

# Authorizer
env GOOS=linux GOARCH=amd64 go build -o /tmp/$AUTHORIZER_NAME lambda.authorizer/main.go
sudo rm /tmp/$AUTHORIZER_NAME && sudo rm /tmp/$AUTHORIZER_NAME.zip
zip -j /tmp/$AUTHORIZER_NAME.zip /tmp/$AUTHORIZER_NAME
aws lambda update-function-code --function-name $AUTHORIZER_NAME --zip-file fileb:///tmp/$AUTHORIZER_NAME.zip

# Login
env GOOS=linux GOARCH=amd64 go build -o /tmp/$LOGIN_NAME lambda.login/main.go
sudo rm /tmp/$LOGIN_NAME && sudo rm /tmp/$LOGIN_NAME.zip
zip -j /tmp/$LOGIN_NAME.zip /tmp/$LOGIN_NAME
aws lambda update-function-code --function-name $LOGIN_NAME --zip-file fileb:///tmp/$LOGIN_NAME.zip

# Change Password
env GOOS=linux GOARCH=amd64 go build -o /tmp/$CHANGE_PASSWORD_NAME lambda.change-password/main.go
sudo rm /tmp/$CHANGE_PASSWORD_NAME && sudo rm /tmp/$CHANGE_PASSWORD_NAME.zip
zip -j /tmp/$CHANGE_PASSWORD_NAME.zip /tmp/$CHANGE_PASSWORD_NAME
aws lambda update-function-code --function-name $CHANGE_PASSWORD_NAME --zip-file fileb:///tmp/$CHANGE_PASSWORD_NAME.zip

# Create User
env GOOS=linux GOARCH=amd64 go build -o /tmp/$CREATE_USER_NAME lambda.create-user/main.go
sudo rm /tmp/$CREATE_USER_NAME && sudo rm /tmp/$CREATE_USER_NAME.zip
zip -j /tmp/$CREATE_USER_NAME.zip /tmp/$CREATE_USER_NAME
aws lambda update-function-code --function-name $CREATE_USER_NAME --zip-file fileb:///tmp/$CREATE_USER_NAME.zip

# Delete User
env GOOS=linux GOARCH=amd64 go build -o /tmp/$DELETE_USER_NAME lambda.delete-user/main.go
sudo rm /tmp/$DELETE_USER_NAME && sudo rm /tmp/$DELETE_USER_NAME.zip
zip -j /tmp/$DELETE_USER_NAME.zip /tmp/$DELETE_USER_NAME
aws lambda update-function-code --function-name $DELETE_USER_NAME --zip-file fileb:///tmp/$DELETE_USER_NAME.zip

# Get User
env GOOS=linux GOARCH=amd64 go build -o /tmp/$GET_USER_NAME lambda.get-user/main.go
sudo rm /tmp/$GET_USER_NAME && sudo rm /tmp/$GET_USER_NAME.zip
zip -j /tmp/$GET_USER_NAME.zip /tmp/$GET_USER_NAME
aws lambda update-function-code --function-name $GET_USER_NAME --zip-file fileb:///tmp/$GET_USER_NAME.zip

# Get Users
env GOOS=linux GOARCH=amd64 go build -o /tmp/$GET_USERS_NAME lambda.get-users/main.go
sudo rm /tmp/$GET_USERS_NAME && sudo rm /tmp/$GET_USERS_NAME.zip
zip -j /tmp/$GET_USERS_NAME.zip /tmp/$GET_USERS_NAME
aws lambda update-function-code --function-name $GET_USERS_NAME --zip-file fileb:///tmp/$GET_USERS_NAME.zip

# Update User
env GOOS=linux GOARCH=amd64 go build -o /tmp/$UPDATE_NAME lambda.update-user/main.go
sudo rm /tmp/$UPDATE_NAME && sudo rm /tmp/$UPDATE_NAME.zip
zip -j /tmp/$UPDATE_NAME.zip /tmp/$UPDATE_NAME
aws lambda update-function-code --function-name $UPDATE_NAME --zip-file fileb:///tmp/$UPDATE_NAME.zip

# Create Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$CREATE_QB_NAME lambda.create-queue-buster/main.go
sudo rm /tmp/$CREATE_QB_NAME && sudo rm /tmp/$CREATE_QB_NAME.zip
zip -j /tmp/$CREATE_QB_NAME.zip /tmp/$CREATE_QB_NAME
aws lambda update-function-code --function-name $CREATE_QB_NAME --zip-file fileb:///tmp/$CREATE_QB_NAME.zip

# Delete Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$DELETE_QB_NAME lambda.delete-queue-buster/main.go
sudo rm /tmp/$DELETE_QB_NAME && sudo rm /tmp/$DELETE_QB_NAME.zip
zip -j /tmp/$DELETE_QB_NAME.zip /tmp/$DELETE_QB_NAME
aws lambda update-function-code --function-name $DELETE_QB_NAME --zip-file fileb:///tmp/$DELETE_QB_NAME.zip

# Get Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$GET_QB_NAME lambda.get-queue-buster/main.go
sudo rm /tmp/$GET_QB_NAME && sudo rm /tmp/$GET_QB_NAME.zip
zip -j /tmp/$GET_QB_NAME.zip /tmp/$GET_QB_NAME
aws lambda update-function-code --function-name $GET_QB_NAME --zip-file fileb:///tmp/$GET_QB_NAME.zip

# Get Queue Busters
env GOOS=linux GOARCH=amd64 go build -o /tmp/$GET_QBS_NAME lambda.get-queue-busters/main.go
sudo rm /tmp/$GET_QBS_NAME && sudo rm /tmp/$GET_QBS_NAME.zip
zip -j /tmp/$GET_QBS_NAME.zip /tmp/$GET_QBS_NAME
aws lambda update-function-code --function-name $GET_QBS_NAME --zip-file fileb:///tmp/$GET_QBS_NAME.zip

# Enable Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$ENABLE_QB_NAME lambda.enable-queue-buster/main.go
sudo rm /tmp/$ENABLE_QB_NAME && sudo rm /tmp/$ENABLE_QB_NAME.zip
zip -j /tmp/$ENABLE_QB_NAME.zip /tmp/$ENABLE_QB_NAME
aws lambda update-function-code --function-name $ENABLE_QB_NAME --zip-file fileb:///tmp/$ENABLE_QB_NAME.zip

# Disable Queue Buster
env GOOS=linux GOARCH=amd64 go build -o /tmp/$DISABLE_QB_NAME lambda.disable-queue-buster/main.go
sudo rm /tmp/$DISABLE_QB_NAME && sudo rm /tmp/$DISABLE_QB_NAME.zip
zip -j /tmp/$DISABLE_QB_NAME.zip /tmp/$DISABLE_QB_NAME
aws lambda update-function-code --function-name $DISABLE_QB_NAME --zip-file fileb:///tmp/$DISABLE_QB_NAME.zip

# Enable Queue Buster Announcements
env GOOS=linux GOARCH=amd64 go build -o /tmp/$ENABLE_ANNOUNCEMENTS_QB_NAME lambda.enable-queue-buster-announcements/main.go
sudo rm /tmp/$ENABLE_ANNOUNCEMENTS_QB_NAME && sudo rm /tmp/$ENABLE_ANNOUNCEMENTS_QB_NAME.zip
zip -j /tmp/$ENABLE_ANNOUNCEMENTS_QB_NAME.zip /tmp/$ENABLE_ANNOUNCEMENTS_QB_NAME
aws lambda update-function-code --function-name $ENABLE_ANNOUNCEMENTS_QB_NAME --zip-file fileb:///tmp/$ENABLE_ANNOUNCEMENTS_QB_NAME.zip

# Enable Queue Buster Announcements
env GOOS=linux GOARCH=amd64 go build -o /tmp/$DISABLE_ANNOUNCEMENTS_QB_NAME lambda.enable-queue-buster-announcements/main.go
sudo rm /tmp/$DISABLE_ANNOUNCEMENTS_QB_NAME && sudo rm /tmp/$DISABLE_ANNOUNCEMENTS_QB_NAME.zip
zip -j /tmp/$DISABLE_ANNOUNCEMENTS_QB_NAME.zip /tmp/$DISABLE_ANNOUNCEMENTS_QB_NAME
aws lambda update-function-code --function-name $DISABLE_ANNOUNCEMENTS_QB_NAME --zip-file fileb:///tmp/$DISABLE_ANNOUNCEMENTS_QB_NAME.zip