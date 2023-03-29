package service

import (
	"be-kelaskita/entity"
	"be-kelaskita/repository"
	"time"
)

type ClassService interface {
	GetClass() ([]entity.Class, error)
	InsertClass(inputUser entity.Class) (entity.Class, error)
	// UpdateClass(inputUser entity.Class, id int) (entity.Class, error)
	// DeleteClass(id int) error
	GetUserByClassId(id int) ([]entity.User, error)
}

type classService struct {
	classRepository repository.ClassRepository
}

func NewClassService(classRepository repository.ClassRepository) *classService {
	return &classService{classRepository}
}

func (c *classService) GetClass() ([]entity.Class, error) {
	classes, err := c.classRepository.GetClass()
	if err != nil {
		return classes, err
	}
	return classes, nil
}

func (c *classService) InsertClass(inputClass entity.Class) (entity.Class, error) {
	var class entity.Class

	class.Name = inputClass.Name
	class.Created_at = time.Now()
	class.Update_at = time.Now()
	class.Teacher_id = inputClass.Teacher_id

	newClass, err := c.classRepository.InsertClass(class)

	if err != nil {
		return newClass, err
	}

	return newClass, nil
}

func (c *classService) GetUserByClassId(id int) ([]entity.User, error) {
	var class entity.Class

	class.ID = id
	users, err := c.classRepository.GetUserByClassId(class)
	if err != nil {
		return users, err
	}
	return users, nil
}
