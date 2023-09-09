package controller

import (
	"encoding/json"
	"errors"
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
	return res.New(w).Status(http.StatusOK).JSON(authors)
}

func (ac *Author) FindOne(w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	author, err := ac.service.FindOne(id)
	if err != nil {
		return res.New(w).NotFoundError("author not found")
	}
	return res.New(w).Status(http.StatusOK).JSON(author)
}

func (ac *Author) Create(w http.ResponseWriter, r *http.Request) error {
	schema, err := ac.parseCreate(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	errMap, hasErr := schema.Validate()
	if hasErr {
		return res.New(w).Status(http.StatusBadRequest).JSON(errMap)
	}
	newAuthor, err := ac.service.Create(schema)
	if err != nil {
		return err
	}
	return res.New(w).Status(http.StatusCreated).JSON(newAuthor)
}

func (ac *Author) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	_, err = ac.service.FindOne(id)
	if err != nil {
		return res.New(w).NotFoundError("author not found")
	}
	schema, err := ac.parseUpdate(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	errMap, hasErr := schema.Validate()
	if hasErr {
		return res.New(w).Status(http.StatusBadRequest).JSON(errMap)
	}
	updatedAuthor, err := ac.service.Update(id, schema)
	if err != nil {
		return err
	}
	return res.New(w).Status(http.StatusOK).JSON(updatedAuthor)
}

func (ac *Author) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	err = ac.service.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (ac *Author) parseCreate(r *http.Request) (*schemas.AuthorCreate, error) {
	defer r.Body.Close()
	schema := new(schemas.AuthorCreate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		return nil, errors.New("error parsing request body")
	}
	return schema, nil
}

func (ac *Author) parseUpdate(r *http.Request) (*schemas.AuthorUpdate, error) {
	defer r.Body.Close()
	schema := new(schemas.AuthorUpdate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		return nil, errors.New("error parsing request body")
	}
	return schema, nil
}
