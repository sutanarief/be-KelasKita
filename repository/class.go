package repository

import (
	"be-kelaskita/entity"
	"database/sql"
)

type ClassRepository interface {
	GetClass() ([]entity.Class, error)
	InsertClass(inputClass entity.Class) (entity.Class, error)
	// UpdateClass(inputclass entity.Class) (entity.Class, error)
	// Deleteclass(class entity.Class) error
	GetUserByClassId(class entity.Class) ([]entity.User, error)
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
			&class.Update_at,
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
		&class.Update_at,
		&class.Teacher_id,
	)

	if err != nil {
		return class, err
	}

	return class, nil
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
