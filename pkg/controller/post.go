package controller

import (
	"encoding/json"
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
	return res.JSON(w, http.StatusOK, posts)
}

func (pc *Post) FindOne(w http.ResponseWriter, r *http.Request) error {
	id := parseId(w, r)
	post, err := pc.service.FindOne(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, post)
}

func (pc *Post) Create(w http.ResponseWriter, r *http.Request) error {
	createSchema := pc.parseCreate(w, r)
	newPost, err := pc.service.Create(createSchema)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, newPost)
}

func (pc *Post) Update(w http.ResponseWriter, r *http.Request) error {
	updateSchema := pc.parseUpdate(w, r)
	updatedPost, err := pc.service.Update(updateSchema)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, updatedPost)
}

func (pc *Post) Delete(w http.ResponseWriter, r *http.Request) error {
	id := parseId(w, r)
	err := pc.service.Delete(id)
	if err != nil {
		res.InternalServerError(w, err.Error())
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (pc *Post) parseCreate(w http.ResponseWriter, r *http.Request) *schemas.PostCreate {
	defer r.Body.Close()
	schema := new(schemas.PostCreate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		res.BadRequest(w, "invalid request body")
	}
	return schema
}

func (pc *Post) parseUpdate(w http.ResponseWriter, r *http.Request) *schemas.PostUpdate {
	defer r.Body.Close()
	schema := new(schemas.PostUpdate)
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		res.BadRequest(w, "invalid request body")
	}
	return schema
}
