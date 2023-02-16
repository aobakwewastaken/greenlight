package data

import (
	"database/sql"
	"errors"
)

type Models struct {
	Movie MovieModel
}

var (
	ErrRecordNotFound = errors.New("record not found")
)

func NewModel(db *sql.DB) Models {
	return Models {
		Movie: MovieModel{DB: db},
	}
}