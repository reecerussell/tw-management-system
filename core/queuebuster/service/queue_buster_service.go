package service

import (
	"fmt"
	"net/http"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/queuebuster/model"
	"github.com/reecerussell/tw-management-system/core/queuebuster/repository"
)

// QueueBusterService is used to provider further validation on the
// domain in which cannot be done at the domain layer.
type QueueBusterService struct {
	repo repository.QueueBusterRepository
}

// NewQueueBusterService returns a new instance of QueueBusterService.
func NewQueueBusterService(repo repository.QueueBusterRepository) *QueueBusterService {
	return &QueueBusterService{
		repo: repo,
	}
}

// EnsureValid ensures the given queue buster is valid and that the
// department doesn't already have a queue buster.
func (s *QueueBusterService) EnsureValid(qb *model.QueueBuster) core.Error {
	_, err := s.repo.Get(qb.GetDepartment())
	if err == nil {
		return core.NewErrorWithStatus(
			fmt.Errorf("the department '%s' already has a queue buster", qb.GetDepartment()),
			http.StatusBadRequest,
		)
	}

	return nil
}
