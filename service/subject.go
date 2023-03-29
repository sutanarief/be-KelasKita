package service

import (
	"be-kelaskita/entity"
	"be-kelaskita/repository"
	"time"
)

type SubjectService interface {
	GetSubject() ([]entity.Subject, error)
	InsertSubject(inputSubject entity.Subject) (entity.Subject, error)
	UpdateSubject(inputSubject entity.Subject, id int) (entity.Subject, error)
	DeleteSubject(id int) error
}

type subjectService struct {
	subjectRepository repository.SubjectRepository
}

func NewSubjectService(subjectRepository repository.SubjectRepository) *subjectService {
	return &subjectService{subjectRepository}
}

func (c *subjectService) GetSubject() ([]entity.Subject, error) {
	subjects, err := c.subjectRepository.GetSubject()
	if err != nil {
		return subjects, err
	}
	return subjects, nil
}

func (c *subjectService) InsertSubject(inputSubject entity.Subject) (entity.Subject, error) {
	var subject entity.Subject

	subject.Name = inputSubject.Name
	subject.Created_at = time.Now()
	subject.Updated_at = time.Now()

	newSubject, err := c.subjectRepository.InsertSubject(subject)

	if err != nil {
		return newSubject, err
	}

	return newSubject, nil
}

func (c *subjectService) UpdateSubject(inputSubject entity.Subject, id int) (entity.Subject, error) {
	var subject entity.Subject

	subject.ID = id

	subject.Name = inputSubject.Name

	subject.Updated_at = time.Now()

	updatedSubject, err := c.subjectRepository.UpdateSubject(subject)
	if err != nil {
		return updatedSubject, err
	}

	return updatedSubject, nil
}

func (c *subjectService) DeleteSubject(id int) error {
	var subject entity.Subject

	subject.ID = id
	err := c.subjectRepository.DeleteSubject(subject)
	if err != nil {
		return err
	}
	return nil
}
