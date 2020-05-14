package model

import (
	"fmt"
	"net/http"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/queuebuster/datamodel"
	"github.com/reecerussell/tw-management-system/core/queuebuster/dto"
)

// QueueBuster is the domain model for the queue buster domain.
type QueueBuster struct {
	department string
	enabled    bool
}

// NewQueueBuster returns a new instance of the QueueBuster domain object.
func NewQueueBuster(d *dto.QueueBuster) *QueueBuster {
	return &QueueBuster{
		department: d.Department,
		enabled:    d.Enabled,
	}
}

// GetDepartment returns the queue buster's department.
func (qb *QueueBuster) GetDepartment() string {
	return qb.department
}

// Enable sets the state of the queue buster to enabled.
//
// If the queue buster is already enabled, an error will be returned.
func (qb *QueueBuster) Enable() core.Error {
	if qb.enabled {
		return core.NewErrorWithStatus(
			fmt.Errorf("the queue buster is already enabled"),
			http.StatusBadRequest,
		)
	}

	qb.enabled = true

	return nil
}

// Disable sets the state of the queue buster to disabled.
//
// If the queue buster is already disabled, an error will be returned.
func (qb *QueueBuster) Disable() core.Error {
	if !qb.enabled {
		return core.NewErrorWithStatus(
			fmt.Errorf("the queue buster is already disabled"),
			http.StatusBadRequest,
		)
	}

	qb.enabled = false

	return nil
}

// DataModel returns a queue buster data model for the QueueBuster.
func (qb *QueueBuster) DataModel() *datamodel.QueueBuster {
	dm := &datamodel.QueueBuster{
		Department: qb.department,
	}

	if qb.enabled {
		dm.Status = "yes"
	} else {
		dm.Status = "no"
	}

	return dm
}

// DTO returns a data transfer object for the QueueBuster.
func (qb *QueueBuster) DTO() *dto.QueueBuster {
	return &dto.QueueBuster{
		Department: qb.department,
		Enabled:    qb.enabled,
	}
}

// QueueBusterFromDataModel returns a QueueBuster domain model for the data model.
func QueueBusterFromDataModel(dm *datamodel.QueueBuster) *QueueBuster {
	qb := &QueueBuster{
		department: dm.Department,
	}

	if dm.Status == "yes" {
		qb.enabled = true
	} else {
		qb.enabled = false
	}

	return qb
}
