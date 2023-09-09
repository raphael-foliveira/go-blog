package schemas

import (
	"time"
)

type AuthorCreate struct {
	Name string `json:"name"`
}

func (ac *AuthorCreate) Validate() (map[string][]string, bool) {
	name := []string{}
	if ac.Name == "" {
		name = append(name, "name cannot be empty")
	} else if len(ac.Name) < 3 {
		name = append(name, "name should at least 3 characters long")
	}
	return map[string][]string{
		"name": name,
	}, len(name) > 0
}

type AuthorUpdate struct {
	Name string `json:"name"`
}

func (au *AuthorUpdate) Validate() (map[string][]string, bool) {
	name := []string{}
	if au.Name == "" {
		name = append(name, "name cannot be empty")
	} else if len(au.Name) < 3 {
		name = append(name, "name should at least 3 characters long")
	}
	return map[string][]string{
		"name": name,
	}, len(name) > 0
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
