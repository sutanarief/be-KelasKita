package service

import (
	"be-kelaskita/entity"
	"be-kelaskita/repository"
	"time"
)

type UserService interface {
	GetUser() ([]entity.User, error)
	InsertUser(inputUser entity.User) (entity.User, error)
	UpdateUser(inputUser entity.User, id int) (entity.User, error)
	DeleteUser(id int) error
	GetUserById(id int) (entity.User, error)
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

	user.Full_name = inputUser.Full_name
	user.Username = inputUser.Username
	user.Password = inputUser.Password
	user.Email = inputUser.Email
	user.Role = inputUser.Role
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.Class_ID = inputUser.Class_ID

	newUser, err := u.userRepository.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (u *userService) UpdateUser(inputUser entity.User, id int) (entity.User, error) {
	var user entity.User

	user.ID = id

	user.Full_name = inputUser.Full_name
	user.Username = inputUser.Username
	user.Password = inputUser.Password
	user.Email = inputUser.Email
	user.Class_ID = inputUser.Class_ID

	user.Updated_at = time.Now()

	updatedUser, err := u.userRepository.UpdateUser(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (u *userService) DeleteUser(id int) error {
	var user entity.User

	user.ID = id
	err := u.userRepository.DeleteUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GetUserById(id int) (entity.User, error) {
	user, err := u.userRepository.GetUserById(id)

	if err != nil {
		return user, err
	}

	return user, nil
}
