package persistence

import (
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/reecerussell/tw-management-system/core/domain/datamodel"
	"github.com/reecerussell/tw-management-system/core/domain/model"
	"github.com/reecerussell/tw-management-system/core/domain/provider"
	"github.com/reecerussell/tw-management-system/core/domain/repository"
)

// Common Errors
var (
	ErrLabelNotFound = errors.New("label not found")
)

// a DynamoDB implementation of repository.LabelRepository
type labelRepository struct {
	client *dynamodb.DynamoDB
	logger *log.Logger
	table string
}

// NewLabelRepository returns a new instance of repository.LabelRepository
// for DynamoDB and the given table.
func NewLabelRepository(table string) repository.LabelRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	logger := log.New(os.Stderr, "[LabelRepository][ERROR]: ", log.LstdFlags)

	return &labelRepository{
		client: dynamodb.New(sess),
		logger: logger,
		table: table,
	}
}

// Get returns a label from DynamoDB.
func (lr *labelRepository) Get(id string) (*model.Label, error) {
	result, err := lr.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(lr.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		lr.logger.Printf("Failed to get label from DynamoDB: %v", err)
		return nil, err
	}

	var lbl datamodel.Label

	err = dynamodbattribute.UnmarshalMap(result.Item, &lbl)
	if err != nil {
		lr.logger.Printf("Failed to unmarshal DynamoDB result: %v", err)
		return nil, err
	}

	if lbl.ID == "" {
		return nil, provider.ErrLabelNotFound
	}

	return model.LabelFromDataModel(&lbl), nil
}

// Create inserts a new label into DynamoDB.
func (lr *labelRepository) Create(l *model.Label) error {
	av, err := dynamodbattribute.MarshalMap(l.DataModel())
	if err != nil {
		lr.logger.Printf("Failed to marshal Label data model: %v", err)
		return err
	}

	in := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(lr.table),
	}

	_, err = lr.client.PutItem(in)
	if err != nil {
		lr.logger.Printf("An error occured while putting the label to DynamoDB: %v", err)
		return err
	}

	return nil
}

// Update updates the label in DynamoDB.
func (lr *labelRepository) Update(l *model.Label) error {
	dm := l.DataModel()
	in := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(dm.Name),
			},
		},
		TableName: aws.String(lr.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(dm.ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set name = :n"),
	}

	_, err := lr.client.UpdateItem(in)
	if err != nil {
		lr.logger.Printf("An error occured while updating label in DynamoDB: %v", err)
		return err
	}

	return nil
}

// Delete removes a label with the given id from the DynamoDB table.
func (lr *labelRepository) Delete(id string) error {
	in := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(lr.table),
	}

	out, err := lr.client.DeleteItem(in)
	if err != nil {
		lr.logger.Printf("An error occured while deleting label in DynamoDB: %v", err)
		return err
	}

	if out.ConsumedCapacity.WriteCapacityUnits != nil &&
		*out.ConsumedCapacity.WriteCapacityUnits < 0 {
		return ErrLabelNotFound
	}

	return nil
}

func (lr *labelRepository) ExistsWithName(name string) (bool, error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(name),
			},
		},
		FilterExpression: aws.String("name = :n"),
		ProjectionExpression: aws.String(dynamodb.SelectCount),
		TableName: aws.String(lr.table),
	}

	result, err := lr.client.Scan(input)
	if err != nil {
		lr.logger.Printf("An error occured while querying DynamoDB: %v", err)
		return false, err
	}

	if result.ScannedCount == nil {
		lr.logger.Printf("Scan count is nil.")
		return false, nil
	}

	return *result.ScannedCount > 0, nil
}