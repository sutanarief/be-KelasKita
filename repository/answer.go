package repository

import (
	"be-kelaskita/entity"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type AnswerRepository interface {
	GetAnswer() ([]entity.Answer, error)
	InsertAnswer(inputAnswer entity.Answer) (entity.Answer, error)
	UpdateAnswer(inputAnswer entity.Answer) (entity.Answer, error)
	DeleteAnswer(answer entity.Answer) error
	GetAnswerById(id int) (entity.Answer, error)
}

type answerRepository struct {
	db *sql.DB
}

func NewAnswerRepository(db *sql.DB) *answerRepository {
	return &answerRepository{db}
}

func (a *answerRepository) GetAnswer() ([]entity.Answer, error) {
	var result []entity.Answer

	sql := "SELECT * FROM answer"
	data, err := a.db.Query(sql)

	if err != nil {
		panic(err)
	}

	for data.Next() {
		var answer entity.Answer

		err := data.Scan(
			&answer.ID,
			&answer.Answer,
			&answer.User_role,
			&answer.Created_at,
			&answer.Updated_at,
			&answer.Task_id,
			&answer.User_id,
		)

		if err != nil {
			panic(err)
		}
		result = append(result, answer)
	}

	return result, nil
}

func (a *answerRepository) InsertAnswer(answer entity.Answer) (entity.Answer, error) {
	sql := `
	INSERT INTO answer (answer, user_role, created_at, updated_at, task_id, user_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING *
	`

	err := a.db.QueryRow(
		sql,
		&answer.Answer,
		&answer.User_role,
		&answer.Created_at,
		&answer.Updated_at,
		&answer.Task_id,
		&answer.User_id,
	).Scan(
		&answer.ID,
		&answer.Answer,
		&answer.User_role,
		&answer.Created_at,
		&answer.Updated_at,
		&answer.Task_id,
		&answer.User_id,
	)

	if err != nil {
		return answer, err
	}

	return answer, nil
}

func (a *answerRepository) UpdateAnswer(answer entity.Answer) (entity.Answer, error) {
	sql := "UPDATE answer SET "
	inputUserValue := reflect.ValueOf(answer)
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
	datas = append(datas, answer.ID)
	sql += " WHERE id = $" + strconv.Itoa(len(datas)) + " RETURNING *"

	err := a.db.QueryRow(
		sql,
		datas...,
	).Scan(
		&answer.ID,
		&answer.Answer,
		&answer.User_role,
		&answer.Created_at,
		&answer.Updated_at,
		&answer.Task_id,
		&answer.User_id,
	)

	if err != nil {
		return answer, err
	}
	return answer, nil
}

func (a *answerRepository) DeleteAnswer(answer entity.Answer) error {
	sql := "DELETE FROM answer WHERE id = $1"
	err := a.db.QueryRow(sql, answer.ID)

	if err != nil {
		return err.Err()
	}

	return nil
}

func (a *answerRepository) GetAnswerById(id int) (entity.Answer, error) {
	var result entity.Answer
	sql := "SELECT * FROM answer WHERE id = $1"
	err := a.db.QueryRow(sql, id).Scan(
		&result.ID,
		&result.Answer,
		&result.User_role,
		&result.Created_at,
		&result.Updated_at,
		&result.Task_id,
		&result.User_id,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}
