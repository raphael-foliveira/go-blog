package models

import "database/sql"

type Author struct {
	Id          int64
	Name        string
	ActiveSince sql.NullTime
}
