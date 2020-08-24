package repository

import "github.com/reecerussell/tw-management-system/core/domain/model"

// LabelRepository is used to handling writing data and commanding
// the data source, following CQRS.
type LabelRepository interface {
	// Get is used to get a Label domain model
	// for the purposes of writing to/with it.
	Get(id string) (*model.Label, error)
	Create(l *model.Label) error
	Update(l *model.Label) error
	Delete(id string) error
	ExistsWithName(name string) (bool, error)
}
