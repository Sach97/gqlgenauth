package db

import (
	"log"

	ctx "github.com/Sach97/ninshoo/context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//DriverSQL holds our driver sql name struct (postgresql, mysql, etc)
type DriverSQL struct {
	Name string
}

//TODO: implements something like this https://dominicstpierre.com/handling-postgresql-and-mongodb-in-one-go-data-package-excerpt-from-my-book-f23bfe8d7cb7 for dealing with mongodb

//OpenDB open a Sql connexion
func (d DriverSQL) OpenDB(config *ctx.Config) (*sqlx.DB, error) {

	log.Println("Loading config db config... ")
	switch {
	case config.DBUrl == "":
		panic("You must set DBUrl env variable")
	}

	log.Println("Database is connecting... ")
	db, err := sqlx.Open(d.Name, config.DBUrl)
	//fmt.Println(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Println("Not connected ")
	}

	log.Println("Database is connected ")
	return db, err
}
