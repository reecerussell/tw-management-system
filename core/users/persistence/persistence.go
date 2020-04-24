package persistence

import (
	"github.com/reecerussell/tw-management-system/core/users/persistence/dynamo"
	"github.com/reecerussell/tw-management-system/core/users/repository"
)

// DynamoDB returns a new instance of repository.UserRepository for DynamoDB.
func DynamoDB() repository.UserRepository {
	return dynamo.New()
}
