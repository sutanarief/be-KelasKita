package entity

import "time"

type Question struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Question   string    `json:"question"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User_role  string    `json:"user_role"`
	Class_id   int       `json:"class_id"`
	User_id    int       `json:"user_id"`
	Subject_id int       `json:"subject_id"`
}
