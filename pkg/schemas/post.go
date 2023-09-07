package schemas

type PostCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostUpdate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Post struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostDetail struct {
	Post
	Author Author `json:"author"`
}
