package schemas

import (
	"errors"
	"time"
)

type PostCreate struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId int64  `json:"author_id"`
}

func (pc *PostCreate) Validate() error {
	if pc.Title == "" {
		return errors.New("post title cannot be empty")
	} else if len(pc.Title) < 3 {
		return errors.New("post title should be at least 3 characters long")
	}
	if pc.Content == "" {
		return errors.New("post content cannot be empty")
	} else if len(pc.Content) < 10 {
		return errors.New("post content should be at least 10 characters long")
	}
	if pc.AuthorId == 0 {
		return errors.New("author_id field cannot be empty")
	}
	return nil
}

type PostUpdate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (pu *PostUpdate) Validate() error {
	if pu.Title == "" {
		return errors.New("post title cannot be empty")
	} else if len(pu.Title) < 3 {
		return errors.New("post title should be at least 3 characters long")
	}
	if pu.Content == "" {
		return errors.New("post content cannot be empty")
	} else if len(pu.Content) < 10 {
		return errors.New("post content should be at least 10 characters long")
	}
	return nil
}

type Post struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostDetail struct {
	Post
	Author Author `json:"author"`
}
