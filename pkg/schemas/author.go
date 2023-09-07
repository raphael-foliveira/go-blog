package schemas

import (
	"time"
)

type AuthorCreate struct {
	Name string `json:"name"`
}

type AuthorUpdate struct {
	Name string `json:"name"`
}

type Author struct {
	Id          int64     `json:"id"`
	ActiveSince time.Time `json:"active_since"`
	Name        string    `json:"name"`
}

type AuthorDetail struct {
	Author
	Posts []Post `json:"posts"`
}
