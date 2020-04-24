package service

import (
	"fmt"
	"net/http"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/users/model"
	"github.com/reecerussell/tw-management-system/core/users/repository"
)

// UserService is a service used for further vaidation, in which is not
// done on the domain level.
type UserService struct {
	repo repository.UserRepository
}

// NewUserService returns a new instance of UserService.
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// EnsureValid validates the user's email and username have not already been taken.
func (s *UserService) EnsureValid(u *model.User) core.Error {
	taken, err := s.repo.ExistsWithUsername(u.GetUsername(), u.GetID())
	if err != nil {
		return err
	}

	if taken {
		verr := fmt.Errorf("username '%s' is already taken", u.GetUsername())
		return core.NewErrorWithStatus(verr, http.StatusBadRequest)
	}

	taken, err = s.repo.ExistsWithEmail(u.GetEmail(), u.GetID())
	if err != nil {
		return err
	}

	if taken {
		verr := fmt.Errorf("email '%s' is already taken", u.GetEmail())
		return core.NewErrorWithStatus(verr, http.StatusBadRequest)
	}

	return nil
}
