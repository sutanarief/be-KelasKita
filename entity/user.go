package entity

import "time"

type User struct {
	ID         int       `json:"id"`
	Full_name  string    `json:"full_name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Class_ID   int       `json:"class_id"`
}
