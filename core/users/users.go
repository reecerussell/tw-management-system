package users

import (
	"github.com/reecerussell/tw-management-system/core/users/persistence"
	"github.com/reecerussell/tw-management-system/core/users/usecase"
)

// Usecase returns a new instance of UserUsecase.
func Usecase() usecase.UserUsecase {
	repo := persistence.DynamoDB()
	return usecase.NewUserUsecase(repo)
}
