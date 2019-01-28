package context

import (
	"fmt"
	"log"

	ctx "github.com/Sach97/gqlgenauth/auth/context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct{}

func (p Postgres) OpenDB(config *ctx.Config) (*sqlx.DB, error) {
	log.Println("Database is connecting... ")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName))
	fmt.Println(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	log.Println("Database is connected ")
	return db, err
}
