package persistence

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/reecerussell/tw-management-system/core/domain/dto"
	"github.com/reecerussell/tw-management-system/core/domain/provider"

)

// a DynamoDB implementation of provider.LabelProvider.
type labelProvider struct {
	client *dynamodb.DynamoDB
	logger *log.Logger
	table string
}

// NewLabelProvider creates a new instance of provider.LabelProvider for DynamoDB.
func NewLabelProvider(table string) provider.LabelProvider {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	logger := log.New(os.Stderr, "[LabelProvider]: ", log.LstdFlags)

	return &labelProvider{
		client: dynamodb.New(sess),
		logger: logger,
		table: table,
	}
}

// Get queries DynamoDB for a label with the given ID.
func (lp *labelProvider) Get(id string) (*dto.Label, error) {
	result, err := lp.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(lp.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		lp.logger.Printf("Failed to get label from DynamoDB: %v", err)
		return nil, err
	}

	var lbl dto.Label

	err = dynamodbattribute.UnmarshalMap(result.Item, &lbl)
	if err != nil {
		lp.logger.Printf("Failed to unmarshal DynamoDB result: %v", err)
		return nil, err
	}

	if lbl.ID == "" {
		return nil, provider.ErrLabelNotFound
	}

	return &lbl, nil
}

// List queries DynamoDB for all labels.
func (lp *labelProvider) List() ([]*dto.Label, error) {
	proj := expression.NamesList(
		expression.Name("id"),
		expression.Name("name"),
	)
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		lp.logger.Printf("An error occured while building the DynamoDB expression: %v", err)
		return nil, err
	}

	result, err := lp.client.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(lp.table),
	})
	if err != nil {
		lp.logger.Printf("An error occured while querying DyanmoDB: %v", err)
		return nil, err
	}

	labels := make([]*dto.Label, *result.Count)

	for i, data := range result.Items {
		var lbl dto.Label

		err = dynamodbattribute.UnmarshalMap(data, &lbl)
		if err != nil {
			lp.logger.Printf("Failed to unmarshal DynamoDB result: %v", err)
			return nil, err
		}

		labels[i] = &lbl
	}

	return labels, nil
}