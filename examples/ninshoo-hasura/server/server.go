package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/Sach97/ninshoo"
	ninshoo_hasura "github.com/Sach97/ninshoo-hasura"
)

const defaultPort = "8081"

func main() {
	userService := ninshoo.NewNinshoo()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/v1alpha1/graphql", handler.GraphQL(ninshoo_hasura.NewExecutableSchema(ninshoo_hasura.Config{Resolvers: &ninshoo_hasura.Resolver{
		UserService: userService,
	}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
