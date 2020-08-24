package usecase

import (
	"fmt"
	"log"
	"os"

	"github.com/reecerussell/tw-management-system/core/domain/dto"
	"github.com/reecerussell/tw-management-system/core/domain/model"
	"github.com/reecerussell/tw-management-system/core/domain/repository"
)

// LabelUsecase is a high-level interface used to manipulate and
// command the label domain.
type LabelUsecase interface {
	Create(d *dto.CreateLabel) (*dto.Label, error)
	Update(d *dto.Label) error
	Delete(id string) error
}

// a basic implementation of the LabelUsecase interface.
type labelUsecase struct {
	repo repository.LabelRepository
	logger *log.Logger
}

// NewLabelUsecase returns a new instance of LabelUsecase with the given repository.
func NewLabelUsecase(repo repository.LabelRepository) LabelUsecase {
	logger := log.New(os.Stderr, "[LabelUsecase]: ", log.LstdFlags)

	return &labelUsecase{
		repo: repo,
		logger: logger,
	}
}

// Create inserts a new label record into the repository.
func (lu *labelUsecase) Create(d *dto.CreateLabel) (*dto.Label, error) {
	l := model.NewLabel(d)

	err := lu.ensureNameIsUnique(l)
	if err != nil {
		return nil, err
	}

	err = lu.repo.Create(l)
	if err != nil {
		lu.logger.Printf("Failed to save the label record: %v", err)
		return nil, err
	}

	return l.DTO(), nil
}

// Update updates an existing label record, then saves it to the repository.
func (lu *labelUsecase) Update(d *dto.Label) error {
	l, err := lu.repo.Get(d.ID)
	if err != nil {
		lu.logger.Printf("Failed to retrieve the label from the repository: %v", err)
		return err
	}

	l.Update(d)
	err = lu.ensureNameIsUnique(l)
	if err != nil {
		return err
	}

	err = lu.repo.Update(l)
	if err != nil {
		lu.logger.Printf("Failed to save the label record: %v", err)
		return err
	}

	return nil
}

func (lu *labelUsecase) ensureNameIsUnique(l *model.Label) error {
	exists, err := lu.repo.ExistsWithName(l.Name())
	if err != nil {
		lu.logger.Printf("Failed to determine whether the label's name exists: %v", err)
		return err
	}

	if exists {
		return fmt.Errorf("a label with name '%s' has already been defined", l.Name())
	}

	return nil
}

// Delete removes the label record from the repository.
func (lu *labelUsecase) Delete(id string) error {
	err := lu.repo.Delete(id)
	if err != nil {
		lu.logger.Printf("Failed to delete label record: %v", err)
		return err
	}

	return nil
}