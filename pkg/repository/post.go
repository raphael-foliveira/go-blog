package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/raphael-foliveira/blog-backend/pkg/models"
)

type Post struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Post {
	return &Post{db}
}

func (pr *Post) Find(conditions ...string) ([]*models.Post, error) {
	query := buildQuery(`SELECT id, title, content, created_at, updated_at, author_id FROM posts`, conditions...)
	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}
	return pr.scanMany(rows)
}

func (pr *Post) FindOne(id int64) (*models.Post, error) {
	row, err := pr.db.Query(`SELECT id, title, content, created_at, updated_at, author_id FROM posts WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return pr.scanOne(row)
}

func (pr *Post) Create(post *models.Post) (*models.Post, error) {
	result, err := pr.db.Query(`
	INSERT INTO posts (title, content, created_at, updated_at, author_id) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, title, content, created_at, updated_at, author_id;
	`,
		post.Title, post.Content, time.Now(), time.Now(), post.AuthorId,
	)
	if err != nil {
		return nil, err
	}
	return pr.scanOne(result)
}

func (pr *Post) Update(post *models.Post) (*models.Post, error) {
	rows, err := pr.db.Query(`
	UPDATE posts SET title = $1, content = $2, updated_at = $3 
	WHERE id = $4
	RETURNING id, title, content, created_at, updated_at, author_id;
	`,
		post.Title, post.Content, time.Now(), post.Id,
	)
	if err != nil {
		return nil, err
	}
	return pr.scanOne(rows)
}

func (pr *Post) Delete(id int64) (int64, error) {
	result, err := pr.db.Exec(`DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}

func (pr *Post) scanOne(row *sql.Rows) (*models.Post, error) {
	for row.Next() {
		return pr.scanSingle(row)
	}
	return nil, errors.New("no rows found")
}

func (pr *Post) scanMany(rows *sql.Rows) ([]*models.Post, error) {
	posts := []*models.Post{}
	for rows.Next() {
		curr, err := pr.scanSingle(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, curr)
	}
	return posts, nil
}

func (pr *Post) scanSingle(rows *sql.Rows) (*models.Post, error) {
	post := models.Post{}
	err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.AuthorId)
	return &post, err
}
