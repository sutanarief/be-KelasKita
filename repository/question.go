package repository

import (
	"be-kelaskita/entity"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type answerList []entity.QuestAns

func (al *answerList) Scan(value interface{}) error {
	// check if the value is nil
	if value == nil {
		*al = nil
		return nil
	}

	// check if the value is a []byte
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("answerList.Scan: unsupported type %T", value)
	}

	// unmarshal the []byte into a slice of Answer objects
	if err := json.Unmarshal(b, al); err != nil {
		return fmt.Errorf("answerList.Scan: %w", err)
	}

	return nil
}

type QuestionRepository interface {
	GetQuestion() ([]entity.Question, error)
	InsertQuestion(inputQuestion entity.Question) (entity.Question, error)
	UpdateQuestion(inputQuestion entity.Question) (entity.Question, error)
	DeleteQuestion(question entity.Question) error
	GetQuestionById(id int) (entity.Question, error)
	GetQuestionWithAnswer(id int) (entity.QuestionWithAns, error)
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
			&question.Question,
			&question.Created_at,
			&question.Updated_at,
			&question.User_role,
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
	INSERT INTO question (title, question, created_at, updated_at,  user_role, class_id, user_id, subject_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING *
	`

	err := q.db.QueryRow(
		sql,
		&question.Title,
		&question.Question,
		&question.Created_at,
		&question.Updated_at,
		&question.User_role,
		&question.Class_id,
		&question.User_id,
		&question.Subject_id,
	).Scan(
		&question.ID,
		&question.Title,
		&question.Question,
		&question.Created_at,
		&question.Updated_at,
		&question.User_role,
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
		&question.Question,
		&question.Created_at,
		&question.Updated_at,
		&question.User_role,
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

func (q *questionRepository) GetQuestionById(id int) (entity.Question, error) {
	var result entity.Question
	sql := "SELECT * FROM question WHERE id = $1"
	err := q.db.QueryRow(sql, id).Scan(
		&result.ID,
		&result.Title,
		&result.Question,
		&result.Created_at,
		&result.Updated_at,
		&result.User_role,
		&result.Class_id,
		&result.User_id,
		&result.Subject_id,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *questionRepository) GetQuestionWithAnswer(id int) (entity.QuestionWithAns, error) {
	var result entity.QuestionWithAns

	sql := `
	SELECT
		q.id,
		q.title,
		q.question,
		q.created_at,
		q.updated_at,
		q.user_role,
		q.class_id,
		q.user_id,
		q.subject_id,
		json_agg(
			json_build_object(
				'id', a.id,
				'answer', a.answer,
				'created_at', a.created_at,
				'updated_at', a.updated_at,
				'user_role', a.user_role,
				'question_id', a.question_id,
				'user_id', a.user_id
			)
		) AS answer
	FROM
		question q
		LEFT JOIN answer a ON q.id = a.question_id
	WHERE q.id = $1
	GROUP BY
		q.id,
		q.title,
		q.question,
		q.created_at,
		q.updated_at,
		q.user_role,
		q.class_id,
		q.user_id,
		q.subject_id`

	err := q.db.QueryRow(sql, id).Scan(
		&result.ID,
		&result.Title,
		&result.Question,
		&result.Created_at,
		&result.Updated_at,
		&result.User_role,
		&result.Class_id,
		&result.User_id,
		&result.Subject_id,
		(*answerList)(&result.Answer),
	)

	if err != nil {
		return result, err
	}

	return result, nil
}
