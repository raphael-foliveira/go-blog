package repository

import (
	"database/sql"
	"errors"

	"github.com/raphael-foliveira/blog-backend/pkg/models"
)

type Author struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *Author {
	return &Author{db}
}

func (ar *Author) Find() ([]*models.Author, error) {
	rows, err := ar.db.Query(`SELECT id, name, active_since FROM authors;`)
	if err != nil {
		return nil, err
	}
	return ar.scanMany(rows)
}

func (ar *Author) FindOne(id int) (*models.Author, error) {
	row, err := ar.db.Query(`SELECT id, name, active_since FROM authors WHERE id = $1;`, id)
	if err != nil {
		return nil, err
	}
	return ar.scanOne(row)
}

func (ar *Author) Create(author *models.Author) (*models.Author, error) {
	result, err := ar.db.Exec(`INSERT INTO authors (name, active_since) VALUES ($1, $2);`, author.Name, author.ActiveSince)
	if err != nil {
		return nil, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return ar.FindOne(int(lastInsertId))
}

func (ar *Author) Update(author models.Author) (*models.Author, error) {
	result, err := ar.db.Exec(`UPDATE authors SET name = $1 WHERE id = $2;`, author.Name, author.Id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected < 1 || err != nil {
		return nil, errors.New("Update failed")
	}
	return ar.FindOne(int(author.Id))
}

func (ar *Author) Delete(id int) (int64, error) {
	result, err := ar.db.Exec(`DELETE FROM authors WHERE id = $1;`, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (ar *Author) scanOne(row *sql.Rows) (*models.Author, error) {
	author := models.Author{}
	err := row.Scan(&author.Id, &author.Name, &author.ActiveSince)
	if err != nil {
		return nil, err
	}
	return &author, err
}

func (ar *Author) scanMany(rows *sql.Rows) ([]*models.Author, error) {
	authors := []*models.Author{}
	for rows.Next() {
		curr, err := ar.scanOne(rows)
		if err != nil {
			return nil, err
		}
		authors = append(authors, curr)
	}
	return authors, nil
}
