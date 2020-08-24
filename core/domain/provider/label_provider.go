package provider

import (
	"errors"
	"github.com/reecerussell/tw-management-system/core/domain/dto"
)

// Common Errors.
var (
	ErrLabelNotFound = errors.New("label was not found")
)

// LabelProvider is used to query different Labels. This
// follows the CQRS pattern by separating queries from commands.
type LabelProvider interface {
	Get(name string) (*dto.Label, error)
	List() ([]*dto.Label, error)
}
