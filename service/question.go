package service

import (
	"be-kelaskita/entity"
	"be-kelaskita/repository"
	"time"
)

type QuestionService interface {
	GetQuestion() ([]entity.Question, error)
	InsertQuestion(inputQuestion entity.Question) (entity.Question, error)
	UpdateQuestion(inputQuestion entity.Question, id int) (entity.Question, error)
	DeleteQuestion(id int) error
	GetQuestionById(id int) (entity.Question, error)
	GetQuestionWithAnswer(id int) (entity.QuestionWithAns, error)
}

type questionService struct {
	questionRepository repository.QuestionRepository
}

func NewQuestionService(questionRepository repository.QuestionRepository) *questionService {
	return &questionService{questionRepository}
}

func (q *questionService) GetQuestion() ([]entity.Question, error) {
	questions, err := q.questionRepository.GetQuestion()

	if err != nil {
		return questions, err
	}

	return questions, nil
}

func (q *questionService) InsertQuestion(inputQuestion entity.Question) (entity.Question, error) {
	var question entity.Question

	question.Title = inputQuestion.Title
	question.Question = inputQuestion.Question
	question.User_role = inputQuestion.User_role
	question.Created_at = time.Now()
	question.Updated_at = time.Now()
	question.Class_id = inputQuestion.Class_id
	question.User_id = inputQuestion.User_id
	question.Subject_id = inputQuestion.Subject_id

	newQuestion, err := q.questionRepository.InsertQuestion(question)
	if err != nil {
		return newQuestion, err
	}

	return newQuestion, nil
}

func (q *questionService) UpdateQuestion(inputQuestion entity.Question, id int) (entity.Question, error) {
	var question entity.Question

	question.ID = id

	question.Title = inputQuestion.Title
	question.Question = inputQuestion.Question
	question.User_role = inputQuestion.User_role

	question.Updated_at = time.Now()
	question.Class_id = inputQuestion.Class_id
	question.User_id = inputQuestion.User_id
	question.Subject_id = inputQuestion.Subject_id

	updatedQuestion, err := q.questionRepository.UpdateQuestion(question)
	if err != nil {
		return updatedQuestion, err
	}

	return updatedQuestion, nil
}

func (q *questionService) DeleteQuestion(id int) error {
	var question entity.Question

	question.ID = id
	err := q.questionRepository.DeleteQuestion(question)
	if err != nil {
		return err
	}
	return nil
}

func (q *questionService) GetQuestionById(id int) (entity.Question, error) {
	question, err := q.questionRepository.GetQuestionById(id)

	if err != nil {
		return question, err
	}

	return question, nil
}

func (q *questionService) GetQuestionWithAnswer(id int) (entity.QuestionWithAns, error) {
	question, err := q.questionRepository.GetQuestionWithAnswer(id)

	if err != nil {
		return question, err
	}

	return question, nil
}
