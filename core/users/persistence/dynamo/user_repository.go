package dynamo

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/users/datamodel"
	"github.com/reecerussell/tw-management-system/core/users/model"
	"github.com/reecerussell/tw-management-system/core/users/repository"
)

// UserRepository is an implementation of the repository.UserRepository
// interface for DynamoDB.
type UserRepository struct {
	client *dynamodb.DynamoDB
	errLog *log.Logger
}

// New returns a new instance of UserRepository.
func New() repository.UserRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &UserRepository{
		client: dynamodb.New(sess),
		errLog: log.New(os.Stderr, "[DYNAMO][ERROR]: ", log.LstdFlags),
	}
}

// Get returns a specific user record from DynamoDB.
func (r *UserRepository) Get(id string) (*model.User, core.Error) {
	result, err := r.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repository.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		r.errLog.Printf("failed to get user from dynamo: %v", err)
		return nil, core.NewError(err)
	}

	var dm datamodel.User

	err = dynamodbattribute.UnmarshalMap(result.Item, &dm)
	if err != nil {
		r.errLog.Printf("failed to read user data: %v", err)
		return nil, core.NewError(err)
	}

	if dm.ID == "" {
		err = fmt.Errorf("no user was found with id: %s", id)
		return nil, core.NewErrorWithStatus(err, http.StatusNotFound)
	}

	return model.UserFromDataModel(&dm), nil
}

// GetByUsername returns a specific user record from DynamoDB with the given username.
func (r *UserRepository) GetByUsername(username string) (*model.User, core.Error) {
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if u.GetUsername() == username {
			return u, nil
		}
	}

	return nil, core.NewErrorWithStatus(
		fmt.Errorf("no user found with username: %s", username),
		http.StatusNotFound,
	)
}

// GetAll returns all user records from DynamoDB.
func (r *UserRepository) GetAll() ([]*model.User, core.Error) {
	proj := expression.NamesList(
		expression.Name("id"),
		expression.Name("username"),
		expression.Name("email"),
		expression.Name("passwordHash"),
	)
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		r.errLog.Printf("failed to build expression: %v", err)
		return nil, core.NewError(err)
	}

	result, err := r.client.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(repository.TableName),
	})
	if err != nil {
		r.errLog.Printf("failed to query users from dynamo: %v", err)
		return nil, core.NewError(err)
	}

	users := make([]*model.User, *result.Count)

	for i, data := range result.Items {
		var dm datamodel.User

		err = dynamodbattribute.UnmarshalMap(data, &dm)
		if err != nil {
			r.errLog.Printf("failed to read user from query: %v", err)
			return nil, core.NewError(err)
		}

		users[i] = model.UserFromDataModel(&dm)
	}

	return users, nil
}

// Add inserts a user record into DynamoDB.
func (r *UserRepository) Add(u *model.User) core.Error {
	av, err := dynamodbattribute.MarshalMap(u.DataModel())
	if err != nil {
		r.errLog.Printf("failed to marshal user data: %v", err)
		return core.NewError(err)
	}

	in := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(repository.TableName),
	}

	_, err = r.client.PutItem(in)
	if err != nil {
		r.errLog.Printf("failed to put user into dynamo: %v", err)
		return core.NewError(err)
	}

	return nil
}

// Update modifies an existing user record in DynamoDB.
func (r *UserRepository) Update(u *model.User) core.Error {
	dm := u.DataModel()
	in := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":u": {
				S: aws.String(dm.Username),
			},
			":e": {
				S: aws.String(dm.Email),
			},
			":p": {
				S: aws.String(dm.PasswordHash),
			},
		},
		TableName: aws.String(repository.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(dm.ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set username = :u, email = :e, passwordHash = :p"),
	}

	_, err := r.client.UpdateItem(in)
	if err != nil {
		r.errLog.Printf("failed to update user: %v", err)
		return core.NewError(err)
	}

	return nil
}

// Delete removes a user record from DynamoDB.
func (r *UserRepository) Delete(id string) core.Error {
	in := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(repository.TableName),
	}

	_, err := r.client.DeleteItem(in)
	if err != nil {
		r.errLog.Printf("failed to delete user: %v", err)
		return core.NewError(err)
	}

	return nil
}

// ExistsWithUsername queries DynamoDB to check if a user without the
// "ignoreID" exists with the given username.
func (r *UserRepository) ExistsWithUsername(username, ignoreID string) (bool, core.Error) {
	users, err := r.GetAll()
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.GetUsername() != username || u.GetID() == ignoreID {
			continue
		}

		return true, nil
	}

	return false, nil
}

// ExistsWithEmail queries DynamoDB to check if a user without the
// "ignoreID" exists with the given email.
func (r *UserRepository) ExistsWithEmail(email, ignoreID string) (bool, core.Error) {
	users, err := r.GetAll()
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.GetEmail() != email || u.GetID() == ignoreID {
			continue
		}

		return true, nil
	}

	return false, nil
}
