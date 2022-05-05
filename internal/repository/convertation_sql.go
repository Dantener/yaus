package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"yaus/internal/entity"
)

type ConvertationSQL struct {
	db *sqlx.DB
}

func NewConvertationSQL(db *sqlx.DB) *ConvertationSQL {
	return &ConvertationSQL{db: db}
}

func (r *ConvertationSQL) CreateShortURL(item entity.Item) (string, error) {
	var shortURl string

	query := fmt.Sprintf(
		"INSERT INTO %s (id, shorturl, originurl) values ($1, $2, $3) RETURNING shorturl", urlsTable,
	)
	row := r.db.QueryRow(query, strconv.Itoa(int(item.Id)), item.ShortURL, item.OriginURL)

	if err := row.Scan(&shortURl); err != nil {
		return "", err
	}

	return shortURl, nil
}

func (r *ConvertationSQL) FindShortURLByOriginURL(item entity.Item) (string, error) {
	var shortURL string

	query := fmt.Sprintf("SELECT shorturl FROM %s WHERE originurl=$1", urlsTable)
	err := r.db.Get(&shortURL, query, item.OriginURL)

	if err != sql.ErrNoRows {
		return shortURL, err
	}

	return shortURL, nil
}

func (r *ConvertationSQL) FindOriginUrlByShortURL(item entity.Item) (string, error) {
	var originURL string

	query := fmt.Sprintf("SELECT ORIGINURL FROM %s WHERE shorturl=$1", urlsTable)
	err := r.db.Get(&originURL, query, item.ShortURL)

	if err != nil {
		return originURL, err
	}

	return originURL, nil
}
