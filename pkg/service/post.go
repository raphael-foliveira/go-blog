package service

import (
	"errors"
	"fmt"

	"github.com/raphael-foliveira/blog-backend/pkg/interfaces"
	"github.com/raphael-foliveira/blog-backend/pkg/models"
	"github.com/raphael-foliveira/blog-backend/pkg/schemas"
)

type Post struct {
	repository       interfaces.Repository[models.Post]
	authorRepository interfaces.Repository[models.Author]
}

func NewPostService(repository interfaces.Repository[models.Post], authorRepository interfaces.Repository[models.Author]) *Post {
	return &Post{repository, authorRepository}
}

func (ps *Post) Find() ([]*schemas.Post, error) {
	posts, err := ps.repository.Find()
	if err != nil {
		return nil, err
	}
	postsDto := []*schemas.Post{}
	for _, post := range posts {
		postsDto = append(postsDto, ps.modelToSchema(post))
	}
	return postsDto, nil
}

func (ps *Post) FindOne(id int64) (*schemas.PostDetail, error) {
	post, err := ps.repository.FindOne(id)
	if err != nil {
		return nil, err
	}
	author, err := ps.authorRepository.FindOne(post.AuthorId)
	if err != nil {
		return nil, err
	}
	return ps.modelToSchemaDetail(post, author), nil
}

func (ps *Post) FindByAuthor(authorId int64) ([]*schemas.Post, error) {
	posts, err := ps.repository.Find(fmt.Sprintf("author_id = %v", authorId))
	if err != nil {
		return nil, err
	}
	postsDto := []*schemas.Post{}
	for _, post := range posts {
		postsDto = append(postsDto, ps.modelToSchema(post))
	}
	return postsDto, nil
}

func (ps *Post) Create(schema *schemas.PostCreate) (*schemas.Post, error) {
	newPost, err := ps.repository.Create(&models.Post{
		Title:    schema.Title,
		Content:  schema.Content,
		AuthorId: schema.AuthorId,
	})
	if err != nil {
		return nil, err
	}
	return &schemas.Post{
		Id:      newPost.Id,
		Title:   newPost.Title,
		Content: newPost.Content,
	}, nil
}

func (ps *Post) Update(id int64, schema *schemas.PostUpdate) (*schemas.Post, error) {
	updatedPost, err := ps.repository.Update(&models.Post{
		Id:      id,
		Title:   schema.Title,
		Content: schema.Content,
	})
	if err != nil {
		return nil, err
	}
	return &schemas.Post{
		Id:      updatedPost.Id,
		Title:   updatedPost.Title,
		Content: updatedPost.Content,
	}, nil
}

func (ps *Post) Delete(id int64) error {
	deleted, err := ps.repository.Delete(id)
	if err != nil {
		return err
	}
	if deleted == 0 {
		return errors.New("post not found")
	}
	return nil
}

func (ps *Post) modelToSchema(model *models.Post) *schemas.Post {
	return &schemas.Post{
		Id:      model.Id,
		Title:   model.Title,
		Content: model.Content,
	}
}

func (ps *Post) modelToSchemaDetail(model *models.Post, author *models.Author) *schemas.PostDetail {
	return &schemas.PostDetail{
		Post: *ps.modelToSchema(model),
		Author: schemas.Author{
			Id:          author.Id,
			Name:        author.Name,
			ActiveSince: author.ActiveSince.Time,
		},
	}
}
