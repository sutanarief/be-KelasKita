package service

import (
	"be-kelaskita/entity"
	"be-kelaskita/repository"
	"time"
)

type UserService interface {
	GetUser() ([]entity.User, error)
	InsertUser(inputUser entity.User) (entity.User, error)
	// UpdateUser(inputUser entity.User, id int) (entity.User, error)
	// DeleteUser(id int) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (u *userService) GetUser() ([]entity.User, error) {
	users, err := u.userRepository.GetUser()

	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *userService) InsertUser(inputUser entity.User) (entity.User, error) {
	var user entity.User

	user.FullName = inputUser.FullName
	user.Username = inputUser.Username
	user.Password = inputUser.Password
	user.Email = inputUser.Email
	user.Role = inputUser.Role
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ClassID = inputUser.ClassID

	newUser, err := u.userRepository.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
