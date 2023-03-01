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
	ErrEditConflit = errors.New("edit conflict")
)

func NewModel(db *sql.DB) Models {
	return Models {
		Movie: MovieModel{DB: db},
	}
}