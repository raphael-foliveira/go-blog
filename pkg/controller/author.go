package controller

import (
	"encoding/json"
	"net/http"

	"github.com/raphael-foliveira/blog-backend/pkg/res"
	"github.com/raphael-foliveira/blog-backend/pkg/schemas"
	"github.com/raphael-foliveira/blog-backend/pkg/service"
)

type Author struct {
	service *service.Author
}

func NewAuthorController(service *service.Author) *Author {
	return &Author{service}
}

func (ac *Author) Find(w http.ResponseWriter, r *http.Request) error {
	authors, err := ac.service.Find()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, authors)
}

func (ac *Author) FindOne(w http.ResponseWriter, r *http.Request) error {
	id := parseId(w, r)
	author, err := ac.service.FindOne(id)
	if err != nil {
		return res.NotFound(w, "author not found")
	}
	return res.JSON(w, http.StatusOK, author)
}

func (ac *Author) Create(w http.ResponseWriter, r *http.Request) error {
	schema := ac.parseCreate(w, r)
	newAuthor, err := ac.service.Create(schema)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, newAuthor)
}

func (ac *Author) Update(w http.ResponseWriter, r *http.Request) error {
	id := parseId(w, r)
	_, err := ac.service.FindOne(id)
	if err != nil {
		return res.NotFound(w, "author not found")
	}
	schema := ac.parseUpdate(w, r)
	updatedAuthor, err := ac.service.Update(id, schema)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, updatedAuthor)
}

func (ac *Author) Delete(w http.ResponseWriter, r *http.Request) error {
	id := parseId(w, r)
	err := ac.service.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (ac *Author) parseCreate(w http.ResponseWriter, r *http.Request) *schemas.AuthorCreate {
	defer r.Body.Close()
	schema := new(schemas.AuthorCreate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		res.BadRequest(w, err.Error())
	}
	return schema
}

func (ac *Author) parseUpdate(w http.ResponseWriter, r *http.Request) *schemas.AuthorUpdate {
	defer r.Body.Close()
	schema := new(schemas.AuthorUpdate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		res.BadRequest(w, err.Error())
	}
	return schema
}
