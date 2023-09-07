package schemas

import "time"

type PostCreate struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId int64  `json:"author_id"`
}

type PostUpdate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
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
