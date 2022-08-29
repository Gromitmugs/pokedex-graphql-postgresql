first:
	echo pokedex-graphql-posgresql

gqlinit:
	go get -u github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen init

gqlgen:
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

run:
	go run server.go