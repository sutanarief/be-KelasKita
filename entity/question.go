package entity

import "time"

type Question struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	User_role   string    `json:"user_role"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Class_id    int       `json:"class_id"`
	User_id     int       `json:"user_id"`
	Subject_id  int       `json:"subject_id"`
}
