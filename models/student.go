package models

import "time"

type Student struct {
	ID 			int 		`json:"id"`
	Identifier  string 		`json:"identifier"`
	Name 		string 		`json:"name"`
	Email 		string 		`json:"email"`
	Semester 	int 		`json:"semester"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time   `json:"updated_at"`
}