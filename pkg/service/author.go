package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/raphael-foliveira/blog-backend/pkg/interfaces"
	"github.com/raphael-foliveira/blog-backend/pkg/models"
	"github.com/raphael-foliveira/blog-backend/pkg/schemas"
)

type Author struct {
	repository  interfaces.Repository[models.Author]
	postService *Post
}

func NewAuthorService(repository interfaces.Repository[models.Author], postService *Post) *Author {
	return &Author{repository, postService}
}

func (as *Author) Find() ([]*schemas.Author, error) {
	authors, err := as.repository.Find()
	if err != nil {
		return nil, err
	}
	authorsDto := []*schemas.Author{}
	for _, author := range authors {
		authorsDto = append(authorsDto, as.modelToSchema(author))
	}
	return authorsDto, nil
}

func (as *Author) FindOne(id int64) (*schemas.AuthorDetail, error) {
	author, err := as.repository.FindOne(id)
	if err != nil {
		return nil, err
	}
	posts, err := as.postService.FindByAuthor(author.Id)

	return as.modelToSchemaDetail(author, posts), nil
}

func (as *Author) Create(authorCreate *schemas.AuthorCreate) (schemas.Author, error) {
	author, err := as.repository.Create(&models.Author{
		Name:        authorCreate.Name,
		ActiveSince: sql.NullTime{Time: time.Now()},
	})
	if err != nil {
		return schemas.Author{}, err
	}
	return schemas.Author{
		Id:          author.Id,
		Name:        author.Name,
		ActiveSince: author.ActiveSince.Time,
	}, nil
}

func (as *Author) Delete(id int64) error {
	deleted, err := as.repository.Delete(id)
	if err != nil {
		return err
	}
	if deleted == 0 {
		return errors.New("author not found")
	}
	return nil
}

func (as *Author) Update(id int64, authorUpdate *schemas.AuthorUpdate) (*schemas.Author, error) {
	modified, err := as.repository.Update(&models.Author{
		Id:   id,
		Name: authorUpdate.Name,
	})
	if err != nil {
		return nil, err
	}
	return &schemas.Author{
		Id:          modified.Id,
		Name:        modified.Name,
		ActiveSince: modified.ActiveSince.Time,
	}, nil
}

func (as *Author) modelToSchema(model *models.Author) *schemas.Author {
	return &schemas.Author{
		Id:          model.Id,
		Name:        model.Name,
		ActiveSince: model.ActiveSince.Time,
	}
}

func (as *Author) modelToSchemaDetail(model *models.Author, posts []schemas.Post) *schemas.AuthorDetail {
	return &schemas.AuthorDetail{
		Author: *as.modelToSchema(model),
		Posts:  posts,
	}
}
