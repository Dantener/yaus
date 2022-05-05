package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const urlsTable = "urls"

type Config struct {
	Host    string
	Port    string
	User    string
	Pass    string
	NameDB  string
	SSLMode string
}

func NewPostgresDB(conf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=%v",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.NameDB, conf.SSLMode,
	),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
