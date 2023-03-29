package repository

import (
	"be-kelaskita/entity"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type UserRepository interface {
	GetUser() ([]entity.User, error)
	InsertUser(inputUser entity.User) (entity.User, error)
	UpdateUser(inputUser entity.User) (entity.User, error)
	// DeleteUser(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) GetUser() ([]entity.User, error) {
	var result []entity.User

	sql := "SELECT * FROM account"
	data, err := u.db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

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
			log.Fatal(err)
		}
		result = append(result, user)
	}

	return result, nil
}

func (u *userRepository) InsertUser(user entity.User) (entity.User, error) {
	sql := `
	INSERT INTO account (full_name, username, password, email, role, created_at, updated_at, class_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING *
	`

	err := u.db.QueryRow(
		sql,
		user.Full_name,
		user.Username,
		user.Password,
		user.Email,
		user.Role,
		user.Created_at,
		user.Updated_at,
		user.Class_ID,
	).Scan(
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
		return user, err
	}

	return user, nil
}

func (u *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	sql := "UPDATE account SET "
	inputUserValue := reflect.ValueOf(user)
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
	datas = append(datas, user.ID)
	sql += " WHERE id = $" + strconv.Itoa(len(datas)) + " RETURNING *"

	err := u.db.QueryRow(
		sql,
		datas...,
	).Scan(
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
		return user, err
	}
	return user, nil
}
