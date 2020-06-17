package usecase

import (
	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/queuebuster/dto"
	"github.com/reecerussell/tw-management-system/core/queuebuster/model"
	"github.com/reecerussell/tw-management-system/core/queuebuster/repository"
	"github.com/reecerussell/tw-management-system/core/queuebuster/service"
)

// QueueBusterUsecase is a usecase used to manage the queue buster domain.
type QueueBusterUsecase interface {
	Get(department string) (*dto.QueueBuster, core.Error)
	GetAll() ([]*dto.QueueBuster, core.Error)
	Create(d *dto.QueueBuster) core.Error
	Enable(department string) core.Error
	Disable(department string) core.Error
	EnableAnnouncements(department string) core.Error
	DisableAnnouncements(department string) core.Error
	Delete(department string) core.Error
}

// queueBusterUsecase is an implementation of QueueBusterUsecase.
type queueBusterUsecase struct {
	serv *service.QueueBusterService
	repo repository.QueueBusterRepository
}

// NewQueueBusterUsecase returns a new instance of QueueBusterUsecase.
func NewQueueBusterUsecase(repo repository.QueueBusterRepository) QueueBusterUsecase {
	return &queueBusterUsecase{
		serv: service.NewQueueBusterService(repo),
		repo: repo,
	}
}

// Get returns a single queue buster for the given department.
func (uc *queueBusterUsecase) Get(department string) (*dto.QueueBuster, core.Error) {
	qb, err := uc.repo.Get(department)
	if err != nil {
		return nil, err
	}

	return qb.DTO(), nil
}

// GetAll returns all queue busters in the data source.
func (uc *queueBusterUsecase) GetAll() ([]*dto.QueueBuster, core.Error) {
	qbs, err := uc.repo.All()
	if err != nil {
		return nil, err
	}

	dtos := make([]*dto.QueueBuster, len(qbs))

	for i, qb := range qbs {
		dtos[i] = qb.DTO()
	}

	return dtos, nil
}

// Create creates a new queue buster record and inserts it into the data source.
func (uc *queueBusterUsecase) Create(d *dto.QueueBuster) core.Error {
	qb := model.NewQueueBuster(d)
	err := uc.serv.EnsureValid(qb)
	if err != nil {
		return err
	}

	err = uc.repo.Add(qb)
	if err != nil {
		return err
	}

	return nil
}

// Enable enables a queue buster.
func (uc *queueBusterUsecase) Enable(department string) core.Error {
	qb, err := uc.repo.Get(department)
	if err != nil {
		return err
	}

	err = qb.Enable()
	if err != nil {
		return err
	}

	err = uc.repo.Update(qb)
	if err != nil {
		return err
	}

	return nil
}

// Disable disables a queue buster.
func (uc *queueBusterUsecase) Disable(department string) core.Error {
	qb, err := uc.repo.Get(department)
	if err != nil {
		return err
	}

	err = qb.Disable()
	if err != nil {
		return err
	}

	err = uc.repo.Update(qb)
	if err != nil {
		return err
	}

	return nil
}

// EnableAnnouncements flags the queue buster with the given department name,
// as having queue announcements enabled.
func (uc *queueBusterUsecase) EnableAnnouncements(department string) core.Error {
	qb, err := uc.repo.Get(department)
	if err != nil {
		return err
	}

	err = qb.EnableAnnouncements()
	if err != nil {
		return err
	}

	err = uc.repo.Update(qb)
	if err != nil {
		return err
	}

	return nil
}

// DisableAnnouncements flags the queue buster with the given department name,
// as having queue announcements disabled.
func (uc *queueBusterUsecase) DisableAnnouncements(department string) core.Error {
	qb, err := uc.repo.Get(department)
	if err != nil {
		return err
	}

	err = qb.DisableAnnouncements()
	if err != nil {
		return err
	}

	err = uc.repo.Update(qb)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a queue buster record from the data source.
func (uc *queueBusterUsecase) Delete(department string) core.Error {
	_, err := uc.repo.Get(department)
	if err != nil {
		return err
	}

	err = uc.repo.Delete(department)
	if err != nil {
		return err
	}

	return nil
}
