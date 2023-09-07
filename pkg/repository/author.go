package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/raphael-foliveira/blog-backend/pkg/models"
)

type Author struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *Author {
	return &Author{db}
}

func (ar *Author) Find(conditions ...string) ([]*models.Author, error) {
	query := buildQuery(`SELECT id, name, active_since FROM authors;`, conditions...)
	rows, err := ar.db.Query(query)
	if err != nil {
		return nil, err
	}
	return ar.scanMany(rows)
}

func (ar *Author) FindOne(id int64) (*models.Author, error) {
	row, err := ar.db.Query(`SELECT id, name, active_since FROM authors WHERE id = $1;`, id)
	if err != nil {
		return nil, err
	}
	return ar.scanOne(row)
}

func (ar *Author) Create(author *models.Author) (*models.Author, error) {
	row, err := ar.db.Query(`INSERT INTO authors (name, active_since) VALUES ($1, $2) RETURNING id, name, active_since;`, author.Name, time.Now())
	if err != nil {
		return nil, err
	}
	return ar.scanOne(row)
}

func (ar *Author) Update(author *models.Author) (*models.Author, error) {
	row, err := ar.db.Query(`UPDATE authors SET name = $1 WHERE id = $2 RETURNING id, name, active_since;`, author.Name, author.Id)
	if err != nil {
		return nil, err
	}
	return ar.scanOne(row)
}

func (ar *Author) Delete(id int64) (int64, error) {
	result, err := ar.db.Exec(`DELETE FROM authors WHERE id = $1;`, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (ar *Author) scanOne(row *sql.Rows) (*models.Author, error) {
	for row.Next() {
		return ar.scanSingle(row)
	}
	return nil, errors.New("no rows")
}

func (ar *Author) scanMany(rows *sql.Rows) ([]*models.Author, error) {
	authors := []*models.Author{}
	for rows.Next() {
		author, err := ar.scanSingle(rows)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (ar *Author) scanSingle(rows *sql.Rows) (*models.Author, error) {
	author := models.Author{}
	err := rows.Scan(&author.Id, &author.Name, &author.ActiveSince)
	return &author, err
}
