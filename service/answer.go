package service

import (
	"be-kelaskita/entity"
	"be-kelaskita/repository"
	"time"
)

type AnswerService interface {
	GetAnswer() ([]entity.Answer, error)
	InsertAnswer(inputAnswer entity.Answer) (entity.Answer, error)
	UpdateAnswer(inputAnswer entity.Answer, id int) (entity.Answer, error)
	DeleteAnswer(id int) error
	GetAnswerById(id int) (entity.Answer, error)
}

type answerService struct {
	answerRepository repository.AnswerRepository
}

func NewAnswerService(answerRepository repository.AnswerRepository) *answerService {
	return &answerService{answerRepository}
}

func (a *answerService) GetAnswer() ([]entity.Answer, error) {
	answer, err := a.answerRepository.GetAnswer()

	if err != nil {
		return answer, err
	}

	return answer, nil
}

func (a *answerService) InsertAnswer(inputAnswer entity.Answer) (entity.Answer, error) {
	var answer entity.Answer

	answer.Answer = inputAnswer.Answer
	answer.User_role = inputAnswer.User_role
	answer.Created_at = time.Now()
	answer.Updated_at = time.Now()
	answer.Task_id = inputAnswer.Task_id
	answer.User_id = inputAnswer.User_id

	newAnswer, err := a.answerRepository.InsertAnswer(answer)
	if err != nil {
		return newAnswer, err
	}

	return newAnswer, nil
}

func (a *answerService) UpdateAnswer(inputAnswer entity.Answer, id int) (entity.Answer, error) {
	var answer entity.Answer

	answer.ID = id

	answer.Answer = inputAnswer.Answer
	answer.User_role = inputAnswer.User_role

	answer.Updated_at = time.Now()
	answer.Task_id = inputAnswer.Task_id
	answer.User_id = inputAnswer.User_id

	updatedAnswer, err := a.answerRepository.UpdateAnswer(answer)
	if err != nil {
		return updatedAnswer, err
	}

	return updatedAnswer, nil
}

func (a *answerService) DeleteAnswer(id int) error {
	var answer entity.Answer

	answer.ID = id
	err := a.answerRepository.DeleteAnswer(answer)
	if err != nil {
		return err
	}
	return nil
}

func (a *answerService) GetAnswerById(id int) (entity.Answer, error) {
	answer, err := a.answerRepository.GetAnswerById(id)

	if err != nil {
		return answer, err
	}

	return answer, nil
}
