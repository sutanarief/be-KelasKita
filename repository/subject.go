package repository

import (
	"be-kelaskita/entity"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type SubjectRepository interface {
	GetSubject() ([]entity.Subject, error)
	InsertSubject(inputSubject entity.Subject) (entity.Subject, error)
	UpdateSubject(inputSubject entity.Subject) (entity.Subject, error)
	DeleteSubject(subject entity.Subject) error
	GetQuestionBySubjectId(subject entity.Subject) ([]entity.Question, error)
}

type subjectRepository struct {
	db *sql.DB
}

func NewSubjectRepository(db *sql.DB) *subjectRepository {
	return &subjectRepository{db}
}

func (s *subjectRepository) GetSubject() ([]entity.Subject, error) {
	var result []entity.Subject

	sql := "SELECT * FROM subject"
	data, err := s.db.Query(sql)

	if err != nil {
		panic(err)
	}

	for data.Next() {
		var subject entity.Subject

		err := data.Scan(
			&subject.ID,
			&subject.Name,
			&subject.Created_at,
			&subject.Updated_at,
		)

		if err != nil {
			panic(err)
		}
		result = append(result, subject)
	}

	return result, nil
}

func (u *subjectRepository) InsertSubject(subject entity.Subject) (entity.Subject, error) {
	sql := `
	INSERT INTO subject (name, created_at, updated_at)
	VALUES ($1, $2, $3)
	RETURNING *
	`

	err := u.db.QueryRow(
		sql,
		subject.Name,
		subject.Created_at,
		subject.Created_at,
	).Scan(
		&subject.ID,
		&subject.Name,
		&subject.Created_at,
		&subject.Updated_at,
	)

	if err != nil {
		return subject, err
	}

	return subject, nil
}

func (s *subjectRepository) UpdateSubject(subject entity.Subject) (entity.Subject, error) {
	sql := "UPDATE subject SET "
	inputUserValue := reflect.ValueOf(subject)
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
	datas = append(datas, subject.ID)
	sql += " WHERE id = $" + strconv.Itoa(len(datas)) + " RETURNING *"

	err := s.db.QueryRow(
		sql,
		datas...,
	).Scan(
		&subject.ID,
		&subject.Name,
		&subject.Created_at,
		&subject.Updated_at,
	)

	if err != nil {
		return subject, err
	}
	return subject, nil
}

func (s *subjectRepository) DeleteSubject(subject entity.Subject) error {
	sql := "DELETE FROM subject WHERE id = $1"
	err := s.db.QueryRow(sql, subject.ID)

	if err != nil {
		return err.Err()
	}

	return nil
}

func (s *subjectRepository) GetQuestionBySubjectId(subject entity.Subject) ([]entity.Question, error) {
	var result []entity.Question

	sql := "SELECT * FROM question WHERE subject_id = $1 ORDER BY created_at DESC"
	data, err := s.db.Query(sql, subject.ID)

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
