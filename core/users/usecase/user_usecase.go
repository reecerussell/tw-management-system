package usecase

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/reecerussell/tw-management-system/core"
	"github.com/reecerussell/tw-management-system/core/jwt"
	"github.com/reecerussell/tw-management-system/core/users/dto"
	"github.com/reecerussell/tw-management-system/core/users/model"
	"github.com/reecerussell/tw-management-system/core/users/repository"
	"github.com/reecerussell/tw-management-system/core/users/service"
)

// DefaultLoginError is a standard error used when login failed. It prevents
// the details of the failure being exposed.
const DefaultLoginError = "Your username and/or password is incorrect."

// UserUsecase is a high-level interface used to manage user records.
type UserUsecase interface {
	Get(id string) (*dto.User, core.Error)
	All() ([]*dto.User, core.Error)
	Create(d *dto.CreateUser) (*dto.User, core.Error)
	Update(d *dto.User) core.Error
	ChangePassword(d *dto.ChangePassword) core.Error
	Delete(id string) core.Error
	Login(d *dto.UserCredentials) (*jwt.AccessToken, core.Error)
}

type userUsecase struct {
	serv *service.UserService
	repo repository.UserRepository
}

// NewUserUsecase returns a new instance of UserUsecase.
func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		serv: service.NewUserService(repo),
		repo: repo,
	}
}

// Get returns a specific user from the datasource.
func (uc *userUsecase) Get(id string) (*dto.User, core.Error) {
	u, err := uc.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return u.DTO(), nil
}

// All returns all users from the datasource.
func (uc *userUsecase) All() ([]*dto.User, core.Error) {
	users, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	dtos := make([]*dto.User, len(users))

	for i, u := range users {
		dtos[i] = u.DTO()
	}

	return dtos, nil
}

// Create creates a user record and inserts into the datasource.
func (uc *userUsecase) Create(d *dto.CreateUser) (*dto.User, core.Error) {
	u, verr := model.NewUser(d)
	if verr != nil {
		return nil, core.NewErrorWithStatus(verr, http.StatusBadRequest)
	}

	err := uc.serv.EnsureValid(u)
	if err != nil {
		return nil, err
	}

	err = uc.repo.Add(u)
	if err != nil {
		return nil, err
	}

	return u.DTO(), nil
}

// Update updates an existing user record.
func (uc *userUsecase) Update(d *dto.User) core.Error {
	u, err := uc.repo.Get(d.ID)
	if err != nil {
		return err
	}

	if err := u.Update(d); err != nil {
		return core.NewErrorWithStatus(err, http.StatusBadRequest)
	}

	err = uc.serv.EnsureValid(u)
	if err != nil {
		return err
	}

	err = uc.repo.Update(u)
	if err != nil {
		return err
	}

	return nil
}

// ChangePassword changes a user's password.
func (uc *userUsecase) ChangePassword(d *dto.ChangePassword) core.Error {
	u, err := uc.repo.Get(d.ID)
	if err != nil {
		return err
	}

	if err := u.ChangePassword(d.CurrentPassword, d.NewPassword); err != nil {
		return core.NewErrorWithStatus(err, http.StatusBadRequest)
	}

	err = uc.repo.Update(u)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a specific user from the datasource.
func (uc *userUsecase) Delete(id string) core.Error {
	_, err := uc.repo.Get(id)
	if err != nil {
		return err
	}

	err = uc.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// Login returns a new JWT access token for a user with matching credentials.
func (uc *userUsecase) Login(d *dto.UserCredentials) (*jwt.AccessToken, core.Error) {
	if d.Username == "" {
		err := fmt.Errorf("username is required")
		return nil, core.NewErrorWithStatus(err, http.StatusBadRequest)
	}

	u, err := uc.repo.GetByUsername(d.Username)
	if err != nil {
		log.Printf("failed to get user: %s", err.Message())
		if err.Status() == http.StatusNotFound {
			return nil, core.NewErrorWithStatus(
				fmt.Errorf(DefaultLoginError),
				http.StatusBadRequest,
			)
		}

		return nil, err
	}

	if verr := u.Verify(d.Password); verr != nil {
		log.Printf("failed to verify user: %s", verr.Error())
		return nil, core.NewErrorWithStatus(
			fmt.Errorf(DefaultLoginError),
			http.StatusBadRequest,
		)
	}

	err = uc.repo.Update(u)
	if err != nil {
		log.Printf("failed to update user: %s", err.Message())
		return nil, core.NewErrorWithStatus(
			fmt.Errorf(DefaultLoginError),
			http.StatusBadRequest,
		)
	}

	secret, terr := core.NewSecret(os.Getenv("AUTH_SECRET_NAME"))
	if terr != nil {
		log.Printf("failed to get secret: %v", terr)
		return nil, core.NewErrorWithStatus(
			fmt.Errorf(DefaultLoginError),
			http.StatusBadRequest,
		)
	}

	pk, serr := secret.RSAPrivateKey("private")
	if serr != nil {
		log.Printf("failed to get private key: %v", err)
		return nil, core.NewErrorWithStatus(
			fmt.Errorf(DefaultLoginError),
			http.StatusBadRequest,
		)
	}

	var claims jwt.Claims
	issued := time.Now().UTC()
	claims.Audiences = []string{os.Getenv("JWT_AUDIENCE")}
	claims.Issuer = os.Getenv("JWT_ISSUER")
	claims.NotBefore = jwt.NewNumericTime(issued)
	claims.Issued = jwt.NewNumericTime(issued)
	claims.Expires = jwt.NewNumericTime(issued.Add(1 * time.Hour))
	claims.Set = map[string]interface{}{
		"user_id": u.GetID(),
	}

	token := jwt.New(&claims)
	if err := token.Sign(pk); err != nil {
		return nil, core.NewError(err)
	}

	return token.AccessToken(), nil
}
