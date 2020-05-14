package queuebuster

import (
	"github.com/reecerussell/tw-management-system/core/queuebuster/persistence"
	"github.com/reecerussell/tw-management-system/core/queuebuster/usecase"
)

// Usecase returns a new usecase for the queue buster domain.
func Usecase() usecase.QueueBusterUsecase {
	repo := persistence.Dynamo()
	return usecase.NewQueueBusterUsecase(repo)
}
