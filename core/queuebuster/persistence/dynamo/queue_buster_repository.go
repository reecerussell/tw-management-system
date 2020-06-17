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
	"github.com/reecerussell/tw-management-system/core/queuebuster/datamodel"
	"github.com/reecerussell/tw-management-system/core/queuebuster/model"
	"github.com/reecerussell/tw-management-system/core/queuebuster/repository"
)

// QueueBusterRepository is an implementation of the repository.QueueBusterRepository
// interface for DynamoDB.
type QueueBusterRepository struct {
	client *dynamodb.DynamoDB
	errLog *log.Logger
}

// New returns a new instance of QueueBusterRepository.
func New() repository.QueueBusterRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &QueueBusterRepository{
		client: dynamodb.New(sess),
		errLog: log.New(os.Stderr, "[DYNAMO][ERROR]: ", log.LstdFlags),
	}
}

// Get returns a single queue buster record for the given department from DynamoDB.
func (r *QueueBusterRepository) Get(department string) (*model.QueueBuster, core.Error) {
	result, err := r.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repository.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"department": {
				S: aws.String(department),
			},
		},
	})
	if err != nil {
		r.errLog.Printf("failed to get queue buster from dynamo: %v", err)
		return nil, core.NewError(err)
	}

	var dm datamodel.QueueBuster

	err = dynamodbattribute.UnmarshalMap(result.Item, &dm)
	if err != nil {
		r.errLog.Printf("failed to read queue buster data: %v", err)
		return nil, core.NewError(err)
	}

	if dm.Department == "" {
		err = fmt.Errorf("no queue buster was found for department: %s", department)
		return nil, core.NewErrorWithStatus(err, http.StatusNotFound)
	}

	return model.QueueBusterFromDataModel(&dm), nil
}

// All retrieves al queue buster records from DynamoDB.
func (r *QueueBusterRepository) All() ([]*model.QueueBuster, core.Error) {
	proj := expression.NamesList(
		expression.Name("department"),
		expression.Name("status"),
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
		r.errLog.Printf("failed to query queue busters from dynamo: %v", err)
		return nil, core.NewError(err)
	}

	busters := make([]*model.QueueBuster, *result.Count)

	for i, data := range result.Items {
		var dm datamodel.QueueBuster

		err = dynamodbattribute.UnmarshalMap(data, &dm)
		if err != nil {
			r.errLog.Printf("failed to read queue buster from query: %v", err)
			return nil, core.NewError(err)
		}

		busters[i] = model.QueueBusterFromDataModel(&dm)
	}

	return busters, nil
}

// Add inserts a new record into DynamoDB.
func (r *QueueBusterRepository) Add(qb *model.QueueBuster) core.Error {
	av, err := dynamodbattribute.MarshalMap(qb.DataModel())
	if err != nil {
		r.errLog.Printf("failed to marshal queue buster data: %v", err)
		return core.NewError(err)
	}

	in := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(repository.TableName),
	}

	_, err = r.client.PutItem(in)
	if err != nil {
		r.errLog.Printf("failed to put queue buster into dynamo: %v", err)
		return core.NewError(err)
	}

	return nil
}

// Update updates a queue buster record in DynamoDB.
func (r *QueueBusterRepository) Update(qb *model.QueueBuster) core.Error {
	dm := qb.DataModel()
	in := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(dm.Status),
			},
			":a": {
				BOOL: aws.Bool(dm.Announcements),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#s": aws.String("status"),
			"#a": aws.String("announcements"),
		},
		TableName: aws.String(repository.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"department": {
				S: aws.String(dm.Department),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #s = :s, #a = :a"),
	}

	_, err := r.client.UpdateItem(in)
	if err != nil {
		r.errLog.Printf("failed to update queue buster: %v", err)
		return core.NewError(err)
	}

	return nil
}

// Delete removes a queue buster records from DynamoDB.
func (r *QueueBusterRepository) Delete(department string) core.Error {
	in := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"department": {
				S: aws.String(department),
			},
		},
		TableName: aws.String(repository.TableName),
	}

	_, err := r.client.DeleteItem(in)
	if err != nil {
		r.errLog.Printf("failed to delete queue buster: %v", err)
		return core.NewError(err)
	}

	return nil
}
