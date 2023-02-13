package models

import "time"

type ToDo struct {
	Id        int64     `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func NewToDo(body string) *ToDo {
	return &ToDo{
		Body:      body,
		CreatedAt: time.Now().Local().UTC(),
	}
}
