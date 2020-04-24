package repository

import (
	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/users/model"
)

// TableName is the name of the DyanmoDB table for users.
const TableName = "users"

// UserRepository is a high-level interface used to manage user data persistence.
type UserRepository interface {
	Get(id string) (*model.User, core.Error)
	GetAll() ([]*model.User, core.Error)
	Add(u *model.User) core.Error
	Update(u *model.User) core.Error
	Delete(id string) core.Error
	ExistsWithUsername(username, ignoreID string) (bool, core.Error)
	ExistsWithEmail(email, ignoreID string) (bool, core.Error)
}
