package persistence

import (
	"github.com/reecerussell/tw-management-system/core/queuebuster/persistence/dynamo"
	"github.com/reecerussell/tw-management-system/core/queuebuster/repository"
)

// Dynamo returns an instance of repository.QueueBusterRepository for DynamoDB.
func Dynamo() repository.QueueBusterRepository {
	return dynamo.New()
}
