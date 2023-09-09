package schemas

import (
	"time"
)

type PostCreate struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId int64  `json:"author_id"`
}

func (pc *PostCreate) Validate() (map[string][]string, bool) {
	title := []string{}
	content := []string{}
	authorId := []string{}
	if pc.Title == "" {
		title = append(title, "post title cannot be empty")
	} else if len(pc.Title) < 3 {
		title = append(title, "post title should be at least 3 characters long")
	}
	if pc.Content == "" {
		content = append(content, "post content cannot be empty")
	} else if len(pc.Content) < 10 {
		content = append(content, "post content should be at least 10 characters long")
	}
	if pc.AuthorId == 0 {
		authorId = append(authorId, "author_id field cannot be empty")
	}
	return map[string][]string{
		"title":     title,
		"content":   content,
		"author_id": authorId,
	}, len(title) > 0 || len(content) > 0 || len(authorId) > 0
}

type PostUpdate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (pu *PostUpdate) Validate() (map[string][]string, bool) {
	title := []string{}
	content := []string{}
	if pu.Title == "" {
		title = append(title, "post title cannot be empty")
	} else if len(pu.Title) < 3 {
		title = append(title, "post title should be at least 3 characters long")
	}
	if pu.Content == "" {
		content = append(content, "post content cannot be empty")
	} else if len(pu.Content) < 10 {
		content = append(content, "post content should be at least 10 characters long")
	}
	return map[string][]string{
		"title":   title,
		"content": content,
	}, len(title) > 0 || len(content) > 0
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
