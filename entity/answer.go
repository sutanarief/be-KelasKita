package entity

import "time"

type Answer struct {
	ID          int       `json:"id"`
	Answer      string    `json:"answer"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	User_role   string    `json:"user_role"`
	Question_id int       `json:"question_id"`
	User_id     int       `json:"user_id"`
}
