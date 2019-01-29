package context

import (
	"fmt"
	"log"

	ctx "github.com/Sach97/gqlgenauth/auth/context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//DriverSQL holds our driver name struct
type DriverSQL struct {
	Name string
}

//TODO: implements something like this https://dominicstpierre.com/handling-postgresql-and-mongodb-in-one-go-data-package-excerpt-from-my-book-f23bfe8d7cb7

//OpenDB open a Sql connexion
func (d DriverSQL) OpenDB(config *ctx.Config) (*sqlx.DB, error) {
	log.Println("Database is connecting... ")
	db, err := sqlx.Open(d.Name, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName))
	fmt.Println(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	log.Println("Database is connected ")
	return db, err
}
