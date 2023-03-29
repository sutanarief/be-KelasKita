package repository

import (
	"be-kelaskita/entity"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type QuestionRepository interface {
	GetQuestion() ([]entity.Question, error)
	InsertQuestion(inputQuestion entity.Question) (entity.Question, error)
	UpdateQuestion(inputQuestion entity.Question) (entity.Question, error)
	DeleteQuestion(question entity.Question) error
}

type questionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *questionRepository {
	return &questionRepository{db}
}

func (q *questionRepository) GetQuestion() ([]entity.Question, error) {
	var result []entity.Question

	sql := "SELECT * FROM question"
	data, err := q.db.Query(sql)

	if err != nil {
		panic(err)
	}

	for data.Next() {
		var question entity.Question

		err := data.Scan(
			&question.ID,
			&question.Title,
			&question.Description,
			&question.User_role,
			&question.Created_at,
			&question.Updated_at,
			&question.Class_id,
			&question.User_id,
			&question.Subject_id,
		)

		if err != nil {
			panic(err)
		}
		result = append(result, question)
	}

	return result, nil
}

func (q *questionRepository) InsertQuestion(question entity.Question) (entity.Question, error) {
	sql := `
	INSERT INTO question (title, description, user_role, created_at, updated_at, class_id, user_id, subject_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING *
	`

	err := q.db.QueryRow(
		sql,
		&question.Title,
		&question.Description,
		&question.User_role,
		&question.Created_at,
		&question.Updated_at,
		&question.Class_id,
		&question.User_id,
		&question.Subject_id,
	).Scan(
		&question.ID,
		&question.Title,
		&question.Description,
		&question.User_role,
		&question.Created_at,
		&question.Updated_at,
		&question.Class_id,
		&question.User_id,
		&question.Subject_id,
	)

	if err != nil {
		return question, err
	}

	return question, nil
}

func (q *questionRepository) UpdateQuestion(question entity.Question) (entity.Question, error) {
	sql := "UPDATE question SET "
	inputUserValue := reflect.ValueOf(question)
	types := inputUserValue.Type()
	index := 1
	var datas []interface{}

	for i := 0; i < inputUserValue.NumField(); i++ {
		if types.Field(i).Name != "Created_at" && types.Field(i).Name != "ID" {
			if !inputUserValue.Field(i).IsZero() {
				sql += fmt.Sprintf("%v = %v, ", strings.ToLower(types.Field(i).Name), "$"+strconv.Itoa(index))
				datas = append(datas, inputUserValue.Field(i).Interface())
				index++
			}
		}
	}

	sql = strings.TrimSuffix(sql, ", ")
	datas = append(datas, question.ID)
	sql += " WHERE id = $" + strconv.Itoa(len(datas)) + " RETURNING *"

	err := q.db.QueryRow(
		sql,
		datas...,
	).Scan(
		&question.ID,
		&question.Title,
		&question.Description,
		&question.User_role,
		&question.Created_at,
		&question.Updated_at,
		&question.Class_id,
		&question.User_id,
		&question.Subject_id,
	)

	if err != nil {
		return question, err
	}
	return question, nil
}

func (q *questionRepository) DeleteQuestion(question entity.Question) error {
	sql := "DELETE FROM question WHERE id = $1"
	err := q.db.QueryRow(sql, question.ID)

	if err != nil {
		return err.Err()
	}

	return nil
}
