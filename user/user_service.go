package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"nurul-iman-blok-m/model"
)

type UserService interface {
	RegisterUser(input RegisterUserInput) (model.User, string, error)
	GetUserByID(ID uint) (model.User, error)
	LoginUser(input LoginUserInput) (model.User, string, error)
}

type userService struct {
	repository UserRepository
}

func NewService(repository UserRepository) *userService {
	return &userService{repository}
}

func (u *userService) RegisterUser(input RegisterUserInput) (model.User, string, error) {
	user := model.User{}
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, "", err
	}
	user.Password = string(passwordHash)
	user.RoleID = input.RoleID

	newUser, errUser := u.repository.SaveUser(user)
	if errUser != nil {
		return newUser, "", errUser
	}

	return newUser, newUser.Role.RoleName, nil
}

func (u *userService) GetUserByID(ID uint) (model.User, error) {
	user, err := u.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found on with that id")
	}
	return user, nil
}

func (u *userService) LoginUser(input LoginUserInput) (model.User, string, error) {
	email := input.Email
	passwordInput := input.Password

	user, err := u.repository.FindByEmail(email)
	if err != nil {
		return user, "", err
	}

	if user.ID == 0 {
		return user, "", errors.New("no User Found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordInput))

	if err != nil {
		return user, "", err
	}

	roleName, _ := u.repository.GetRoleForResponse(user)

	return user, roleName.Role.RoleName, nil
}
