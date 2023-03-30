package entity

import "time"

type Answer struct {
	ID         int       `json:"id"`
	Answer     string    `json:"answer"`
	User_role  string    `json:"user_role"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Task_id    int       `json:"task_id"`
	User_id    int       `json:"user_id"`
}
