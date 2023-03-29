package entity

import "time"

type Class struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Update_at  time.Time `json:"updated_at"`
	Teacher_id int       `json:"teacher_id"`
}
