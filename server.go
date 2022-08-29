package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Gromitmugs/pokedex-graphql-sqlite/graph"
	"github.com/Gromitmugs/pokedex-graphql-sqlite/graph/generated"
	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const defaultPort = "8001"

func main() {
	r := chi.NewRouter()
	serviceName := "localhost"
	if env, ok := os.LookupEnv("SERVICE"); ok {
		serviceName = env
	}
	dsn := "postgres://user:12345@" + serviceName + ":5432/?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, sqlitedialect.New())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: graph.Database{
			Postgres: db,
		},
	}}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf(dsn)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
