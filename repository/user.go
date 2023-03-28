package repository

import (
	"be-kelaskita/entity"
	"database/sql"
	"log"
)

type UserRepository interface {
	GetUser() ([]entity.User, error)
	InsertUser(inputUser entity.User) (entity.User, error)
	// UpdateUser(inputUser entity.User, id int) (entity.User, error)
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
			&user.FullName,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.ClassID,
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
		user.FullName,
		user.Username,
		user.Password,
		user.Email,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
		user.ClassID,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.ClassID,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}
