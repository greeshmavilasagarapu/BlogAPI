package main

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	Author_id  uuid.UUID `json:"author_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Author struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}
