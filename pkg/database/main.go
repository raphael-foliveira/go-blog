package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func MustConnect() *sql.DB {
	var (
		port     = os.Getenv("POSTGRES_PORT")
		name     = os.Getenv("POSTGRES_DB")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		host     = os.Getenv("POSTGRES_HOST")
	)
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		host, port, user, password, name,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	_, err = createSchema(db)
	if err != nil {
		panic(err)
	}
	return db
}

func createSchema(db *sql.DB) (sql.Result, error) {
	return db.Exec(`
		CREATE TABLE IF NOT EXISTS authors (
			id serial primary key,
			name varchar not null,
			active_since timestamp not null default now()
		);

		CREATE TABLE IF NOT EXISTS posts (
			id serial primary key,
			title varchar not null,
			content text not null,
			created_at timestamp not null default now(),
			updated_at timestamp not null default now(),
			author_id integer,
			CONSTRAINT fk_author
			FOREIGN KEY (author_id) REFERENCES authors(id)
		)
	`)
}
