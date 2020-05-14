package repository

import (
	"os"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/queuebuster/model"
)

// TableName is the table name for the queue buster records.
var TableName = os.Getenv("QUEUE_BUSTER_TABLE_NAME")

// QueueBusterRepository is a high-level interface used to read and
// write queue buster records to a data source.
type QueueBusterRepository interface {
	Get(department string) (*model.QueueBuster, core.Error)
	All() ([]*model.QueueBuster, core.Error)
	Add(qb *model.QueueBuster) core.Error
	Update(qb *model.QueueBuster) core.Error
	Delete(department string) core.Error
}
