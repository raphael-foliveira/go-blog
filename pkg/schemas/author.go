package schemas

import (
	"errors"
	"time"
)

type AuthorCreate struct {
	Name string `json:"name"`
}

func (ac *AuthorCreate) Validate() error {
	if ac.Name == "" {
		return errors.New("name cannot be empty")
	}
	if len(ac.Name) < 3 {
		return errors.New("name should at least 3 characters long")
	}
	return nil
}

type AuthorUpdate struct {
	Name string `json:"name"`
}

func (au *AuthorUpdate) Validate() error {
	if au.Name == "" {
		return errors.New("name cannot be empty")
	}
	if len(au.Name) < 3 {
		return errors.New("name should at least 3 characters long")
	}
	return nil
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
