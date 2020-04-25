package model

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"

	"github.com/reecerussell/tw-management-system/core/users/dto"

	"github.com/reecerussell/tw-management-system/core/users/datamodel"

	"golang.org/x/crypto/pbkdf2"
)

// User variables.
const (
	MinimumUsernameLength  = 5
	MaximumUsernameLength  = 50
	PasswordSaltSize       = 64
	PasswordKeySize        = 256
	PasswordHashIterations = 15000
	MinimumPasswordLength  = 6
)

// Common domain errors.
var (
	ErrUserNotFound = errors.New("user not found")

	ErrUsernameRequired = errors.New("username is required")
	ErrUsernameTooShort = fmt.Errorf("username must be greater than %d characters long", MinimumUsernameLength-1)
	ErrUsernameTooLong  = fmt.Errorf("username cannot be greater than %d characters long", MaximumUsernameLength)
	ErrUsernameInvalid  = errors.New("username must start with a letter or number and can only contain letters, numbers, underscores and dots")

	ErrEmailRequired = errors.New("email is required")
	ErrEmailInvalid  = errors.New("email is invalid")

	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordTooShort = fmt.Errorf("password must be greater than %d characters", MinimumPasswordLength-1)
	ErrPasswordInvalid  = errors.New("password is invalid")
)

var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9]+([._]?[a-zA-Z0-9]+)*$")
var emailRegexp = regexp.MustCompile("[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,6}")
var hashAlg = sha256.New

type User struct {
	id           string
	username     string
	email        string
	passwordHash string
}

// NewUser creates a new User domain object, ensuring the data is valid.
func NewUser(d *dto.CreateUser) (*User, error) {
	u := &User{
		id: uuid.New().String(),
	}

	err := u.UpdateUsername(d.Username)
	if err != nil {
		return nil, err
	}

	err = u.UpdateEmail(d.Email)
	if err != nil {
		return nil, err
	}

	err = u.setPassword(d.Password)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetID returns the user's id.
func (u *User) GetID() string {
	return u.id
}

// GetUsername returns the user's username.
func (u *User) GetUsername() string {
	return u.username
}

// GetEmail returns the user's email.
func (u *User) GetEmail() string {
	return u.email
}

// Update updates the user's properties.
func (u *User) Update(d *dto.User) error {
	err := u.UpdateUsername(d.Username)
	if err != nil {
		return err
	}

	err = u.UpdateEmail(d.Email)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUsername updates the user's username, ensuring it is valid.
func (u *User) UpdateUsername(username string) error {
	l := len(username)
	switch true {
	case username == "":
		return ErrUsernameRequired
	case l < MinimumUsernameLength:
		return ErrUsernameTooShort
	case l > MaximumUsernameLength:
		return ErrUsernameTooLong
	case !usernameRegex.MatchString(username):
		return ErrUsernameInvalid
	}

	u.username = username

	return nil
}

// UpdateEmail updates a user's email address ensuring it's valid.
func (u *User) UpdateEmail(email string) error {
	switch true {
	case email == "":
		return ErrEmailRequired
	case !emailRegexp.MatchString(email):
		return ErrEmailInvalid
	}

	u.email = email

	return nil
}

// ChangePassword changes a user's password once ensuring the current
// password is matches and the new password is valid,
func (u *User) ChangePassword(current, new string) error {
	if !verify(current, u.passwordHash) {
		return ErrPasswordInvalid
	}

	return u.setPassword(new)
}

// Verify is used to verify a given password against the user's hash.
//
// If the given password is equal to the user's current hash, the user's
// password will be rehashed.
func (u *User) Verify(password string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}

	if password == u.passwordHash {
		return u.setPassword(password)
	}

	if !verify(password, u.passwordHash) {
		return fmt.Errorf("password is invalid")
	}

	return nil
}

func (u *User) setPassword(password string) error {
	switch true {
	case password == "":
		return ErrPasswordRequired
	case len(password) < MinimumPasswordLength:
		return ErrPasswordTooShort
	}

	u.passwordHash = hash(password)

	return nil
}

func hash(pwd string) string {
	salt := make([]byte, PasswordSaltSize)
	rand.Read(salt)

	salt64 := base64.StdEncoding.EncodeToString(salt)
	salt = bytes.NewBufferString(salt64).Bytes()

	df := pbkdf2.Key([]byte(pwd), salt, PasswordHashIterations, PasswordKeySize, hashAlg)
	cipher := base64.StdEncoding.EncodeToString(df)

	return fmt.Sprintf("%s.%s", cipher, salt64)
}

func verify(pwd, hash string) bool {
	if pwd == "" || !strings.Contains(hash, ".") {
		return false
	}

	parts := strings.Split(hash, ".")
	cipher, salt64 := parts[0], parts[1]

	salt := bytes.NewBufferString(salt64).Bytes()
	df := pbkdf2.Key([]byte(pwd), salt, PasswordHashIterations, PasswordKeySize, hashAlg)

	return base64.StdEncoding.EncodeToString(df) == cipher
}

// DataModel returns an instance of datamodel.User for the user.
func (u *User) DataModel() *datamodel.User {
	return &datamodel.User{
		ID:           u.id,
		Username:     u.username,
		Email:        u.email,
		PasswordHash: u.passwordHash,
	}
}

// DTO returns a dto.User for the user.
func (u *User) DTO() *dto.User {
	return &dto.User{
		ID:       u.id,
		Username: u.username,
		Email:    u.email,
	}
}

// UserFromDataModel returns a new instance of User with data from
// the given data model. This method assumes the data is valid and
// comes from a data source.
func UserFromDataModel(dm *datamodel.User) *User {
	return &User{
		id:           dm.ID,
		username:     dm.Username,
		email:        dm.Email,
		passwordHash: dm.PasswordHash,
	}
}
