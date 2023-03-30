package repository

import (
	"be-kelaskita/entity"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ClassRepository interface {
	GetClass() ([]entity.Class, error)
	InsertClass(inputClass entity.Class) (entity.Class, error)
	UpdateClass(inputclass entity.Class) (entity.Class, error)
	DeleteClass(class entity.Class) error
	GetUserByClassId(class entity.Class) ([]entity.User, error)
	GetQuestionByClassId(class entity.Class) ([]entity.Question, error)
}

type classRepository struct {
	db *sql.DB
}

func NewClassRepository(db *sql.DB) *classRepository {
	return &classRepository{db}
}

func (c *classRepository) GetClass() ([]entity.Class, error) {
	var result []entity.Class

	sql := "SELECT * FROM class"
	data, err := c.db.Query(sql)

	if err != nil {
		panic(err)
	}

	for data.Next() {
		var class entity.Class

		err := data.Scan(
			&class.ID,
			&class.Name,
			&class.Created_at,
			&class.Updated_at,
			&class.Teacher_id,
		)

		if err != nil {
			panic(err)
		}
		result = append(result, class)
	}

	return result, nil
}

func (u *classRepository) InsertClass(class entity.Class) (entity.Class, error) {
	sql := `
	INSERT INTO class (name, created_at, updated_at, teacher_id)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	err := u.db.QueryRow(
		sql,
		class.Name,
		class.Created_at,
		class.Created_at,
		class.Teacher_id,
	).Scan(
		&class.ID,
		&class.Name,
		&class.Created_at,
		&class.Updated_at,
		&class.Teacher_id,
	)

	if err != nil {
		return class, err
	}

	return class, nil
}

func (c *classRepository) UpdateClass(class entity.Class) (entity.Class, error) {
	sql := "UPDATE class SET "
	inputUserValue := reflect.ValueOf(class)
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
	datas = append(datas, class.ID)
	sql += " WHERE id = $" + strconv.Itoa(len(datas)) + " RETURNING *"

	err := c.db.QueryRow(
		sql,
		datas...,
	).Scan(
		&class.ID,
		&class.Name,
		&class.Created_at,
		&class.Updated_at,
		&class.Teacher_id,
	)

	if err != nil {
		return class, err
	}
	return class, nil
}

func (c *classRepository) DeleteClass(class entity.Class) error {
	sql := "DELETE FROM class WHERE id = $1"
	err := c.db.QueryRow(sql, class.ID)

	if err != nil {
		return err.Err()
	}

	return nil
}

func (c *classRepository) GetUserByClassId(class entity.Class) ([]entity.User, error) {
	var result []entity.User

	sql := "SELECT * FROM account WHERE class_id = $1 ORDER BY role DESC"
	data, err := c.db.Query(sql, class.ID)

	if err != nil {
		panic(err)
	}

	for data.Next() {
		var user entity.User

		err := data.Scan(
			&user.ID,
			&user.Full_name,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Role,
			&user.Created_at,
			&user.Updated_at,
			&user.Class_ID,
		)

		if err != nil {
			panic(err)
		}
		result = append(result, user)
	}

	return result, nil
}

func (c *classRepository) GetQuestionByClassId(class entity.Class) ([]entity.Question, error) {
	var result []entity.Question

	sql := "SELECT * FROM question WHERE class_id = $1 ORDER BY created_at DESC"
	data, err := c.db.Query(sql, class.ID)

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
