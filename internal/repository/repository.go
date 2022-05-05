package repository

import (
	"github.com/jmoiron/sqlx"
	"yaus/internal/entity"
)

type Convertation interface {
	CreateShortURL(url entity.Item) (string, error)
	FindShortURLByOriginURL(url entity.Item) (string, error)
	FindOriginUrlByShortURL(item entity.Item) (string, error)
}

type Repository struct {
	Convertation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Convertation: NewConvertationSQL(db),
	}
}
