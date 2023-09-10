package controller

import (
	"encoding/json"
	"errors"
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
	return res.New(w).Status(http.StatusOK).JSON(posts)
}

func (pc *Post) FindOne(w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	post, err := pc.service.FindOne(id)
	if err != nil {
		return err
	}
	return res.New(w).Status(http.StatusOK).JSON(post)
}

func (pc *Post) Create(w http.ResponseWriter, r *http.Request) error {
	schema, err := pc.parseCreate(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	errMap, hasErr := schema.Validate()
	if hasErr {
		return res.New(w).Status(http.StatusBadRequest).JSON(errMap)
	}
	newPost, err := pc.service.Create(schema)
	if err != nil {
		return err
	}
	return res.New(w).Status(http.StatusCreated).JSON(newPost)
}

func (pc *Post) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	schema, err := pc.parseUpdate(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	errMap, hasErr := schema.Validate()
	if hasErr {
		return res.New(w).Status(http.StatusBadRequest).JSON(errMap)
	}
	updatedPost, err := pc.service.Update(id, schema)
	if err != nil {
		return err
	}
	return res.New(w).Status(http.StatusOK).JSON(updatedPost)
}

func (pc *Post) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r)
	if err != nil {
		return res.New(w).BadRequestError(err.Error())
	}
	err = pc.service.Delete(id)
	if err != nil {
		res.New(w).InternalServerError(err.Error())
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (pc *Post) parseCreate(r *http.Request) (*schemas.PostCreate, error) {
	defer r.Body.Close()
	schema := new(schemas.PostCreate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		return nil, errors.New("error parsing request body")
	}
	return schema, nil
}

func (pc *Post) parseUpdate(r *http.Request) (*schemas.PostUpdate, error) {
	defer r.Body.Close()
	schema := new(schemas.PostUpdate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		return nil, errors.New("error parsing request body")
	}
	return schema, nil
}
