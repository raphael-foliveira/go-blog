package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/raphael-foliveira/blog-backend/pkg/res"
	"github.com/raphael-foliveira/blog-backend/pkg/schemas"
	"github.com/raphael-foliveira/blog-backend/pkg/service"
)

type Post struct {
	service *service.Post
}

func NewPostController(service *service.Post) *Post {
	return &Post{service}
}

func (pc *Post) Find(w http.ResponseWriter, r *http.Request) error {
	posts, err := pc.service.Find()
	if err != nil {
		return err
	}
	return res.New(w, http.StatusOK, posts)
}

func (pc *Post) FindOne(w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r)
	if err != nil {
		return err
	}
	post, err := pc.service.FindOne(id)
	if err != nil {
		return err
	}
	return res.New(w, http.StatusOK, post)
}

func (pc *Post) Create(w http.ResponseWriter, r *http.Request) error {
	schema, err := pc.parseCreate(r)
	if err != nil {
		return err
	}
	errMap, hasErr := schema.Validate()
	if hasErr {
		return fmt.Errorf("%v", errMap)
	}
	newPost, err := pc.service.Create(schema)
	if err != nil {
		return err
	}
	return res.New(w, http.StatusCreated, newPost)
}

func (pc *Post) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r)
	if err != nil {
		return err
	}
	schema, err := pc.parseUpdate(r)
	if err != nil {
		return err
	}
	errMap, hasErr := schema.Validate()
	if hasErr {
		return fmt.Errorf("%v", errMap)
	}
	updatedPost, err := pc.service.Update(id, schema)
	if err != nil {
		return err
	}
	return res.New(w, http.StatusOK, updatedPost)
}

func (pc *Post) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r)
	if err != nil {
		return err
	}
	err = pc.service.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (pc *Post) parseCreate(r *http.Request) (*schemas.PostCreate, error) {
	defer r.Body.Close()
	schema := new(schemas.PostCreate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		return nil, ErrParsingRequestBody
	}
	return schema, nil
}

func (pc *Post) parseUpdate(r *http.Request) (*schemas.PostUpdate, error) {
	defer r.Body.Close()
	schema := new(schemas.PostUpdate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		return nil, ErrParsingRequestBody
	}
	return schema, nil
}

var ErrParsingRequestBody = errors.New("error parsing request body")
