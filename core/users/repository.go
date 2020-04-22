package users

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// TableName is the name of the table in DynamoDB.
const TableName = "users"

type Repository interface {
	Get(id string) (*User, error)
}

type repository struct {
	svc    *dynamodb.DynamoDB
	stderr *log.Logger
}

func NewRepository() Repository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &repository{
		svc:    dynamodb.New(sess),
		stderr: log.New(os.Stderr, "[REPOSITORY][ERROR]: ", log.LstdFlags),
	}
}

func (r *repository) Get(id string) (*User, error) {
	result, err := r.svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		r.stderr.Printf("get item: %v\n", err)
		return nil, fmt.Errorf("failed to get user from dynamo")
	}

	var user User

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		r.stderr.Printf("unmarshal: %v\n", err)
		return nil, fmt.Errorf("failed to read user data from dynamo")
	}

	if user.ID == "" {
		r.stderr.Printf("not found: no user exists with id '%s'.\n", id)
		return nil, fmt.Errorf("user could not be found")
	}

	return &user, nil
}
