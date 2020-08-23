package model

import (
	"github.com/google/uuid"
	"github.com/reecerussell/tw-management-system/core/domain/datamodel"
	"github.com/reecerussell/tw-management-system/core/domain/dto"
)

// Label is a domain object for the hours of operations labels.
type Label struct {
	id   string
	name string
}

func NewLabel(cl *dto.CreateLabel) *Label {
	l := &Label{
		id: uuid.New().String(),
	}

	l.updateName(cl.Name)

	return l
}

// Update updates the label's values, such as name.
func (l *Label) Update(d *dto.Label) {
	l.updateName(d.Name)
}

// updateName is used to update the label's name value.
//
// There are currently no business rules in what a name
// can be, so this method simply sets the value of the label's name.
func (l *Label) updateName(name string) {
	l.name = name
}

// DataModel returns a *datamodel.Label for the label.
func (l *Label) DataModel() *datamodel.Label {
	return &datamodel.Label{
		ID: l.id,
		Name: l.name,
	}
}

// LabelFromDataModel is a function used to return a Label
// domain model, without going through validation like NewLabel.
//
// This is to be used by the repository layer only.
func LabelFromDataModel(dm *datamodel.Label) *Label {
	return &Label {
		id: dm.ID,
		name: dm.Name,
	}
}