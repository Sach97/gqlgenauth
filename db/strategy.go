package db

import (
	ctx "github.com/Sach97/ninshoo/context"
	"github.com/jmoiron/sqlx"
)

//Strategy represents our strategy interface
type Strategy interface {
	OpenDB(config *ctx.Config) (*sqlx.DB, error)
}

//DB holds our DB strategy
type DB struct {
	Strategy Strategy
}

//OpenDB is the strategy method for OpenDB
func (s *DB) OpenDB(config *ctx.Config) (*sqlx.DB, error) {
	db, err := s.Strategy.OpenDB(config)
	return db, err
}
