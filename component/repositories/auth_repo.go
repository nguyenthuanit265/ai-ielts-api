package repositories

import "github.com/jmoiron/sqlx"

type AuthRepo interface {
}
type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) AuthRepo {
	return &authRepo{
		db: db,
	}
}
