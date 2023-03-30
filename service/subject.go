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
	GetQuestionBySubjectId(id int) ([]entity.Question, error)
}

type subjectService struct {
	subjectRepository repository.SubjectRepository
}

func NewSubjectService(subjectRepository repository.SubjectRepository) *subjectService {
	return &subjectService{subjectRepository}
}

func (s *subjectService) GetSubject() ([]entity.Subject, error) {
	subjects, err := s.subjectRepository.GetSubject()
	if err != nil {
		return subjects, err
	}
	return subjects, nil
}

func (s *subjectService) InsertSubject(inputSubject entity.Subject) (entity.Subject, error) {
	var subject entity.Subject

	subject.Name = inputSubject.Name
	subject.Created_at = time.Now()
	subject.Updated_at = time.Now()

	newSubject, err := s.subjectRepository.InsertSubject(subject)

	if err != nil {
		return newSubject, err
	}

	return newSubject, nil
}

func (s *subjectService) UpdateSubject(inputSubject entity.Subject, id int) (entity.Subject, error) {
	var subject entity.Subject

	subject.ID = id

	subject.Name = inputSubject.Name

	subject.Updated_at = time.Now()

	updatedSubject, err := s.subjectRepository.UpdateSubject(subject)
	if err != nil {
		return updatedSubject, err
	}

	return updatedSubject, nil
}

func (s *subjectService) DeleteSubject(id int) error {
	var subject entity.Subject

	subject.ID = id
	err := s.subjectRepository.DeleteSubject(subject)
	if err != nil {
		return err
	}
	return nil
}

func (s *subjectService) GetQuestionBySubjectId(id int) ([]entity.Question, error) {
	var subject entity.Subject

	subject.ID = id
	questions, err := s.subjectRepository.GetQuestionBySubjectId(subject)
	if err != nil {
		return questions, err
	}
	return questions, nil
}
