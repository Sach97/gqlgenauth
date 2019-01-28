package context

import (
	ctx "github.com/Sach97/gqlgenauth/auth/context"
	"github.com/jmoiron/sqlx"
)

type Strategy interface {
	OpenDB(config *ctx.Config) (*sqlx.DB, error)
}

type DB struct {
	Strategy Strategy
}

func (o *DB) OpenDB(config *ctx.Config) (*sqlx.DB, error) {
	db, err := o.Strategy.OpenDB(config)
	return db, err
}
