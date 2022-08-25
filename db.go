package main

// https://github.com/uptrace/bun/blob/master/example/basic/main.go

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func main2() {
	sqlite, err := sql.Open(sqliteshim.ShimName, "pokedex.db")
	if err != nil {
		panic(err)
	}
	sqlite.SetMaxOpenConns(1)

	db := bun.NewDB(sqlite, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	// List All
	// pokedex := make([]Pokemon, 0)
	// if err := db.NewSelect().Model(&pokedex).OrderExpr("id ASC").Scan(ctx); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf(pokedex[0].Name)

	// Insert
	// pokemon := &Pokemon{Name: "Charmander", Description: "It has a preference for hot things.", Category: "Lizard", Type: `{"type1":"Fire"}`, Abilities: `{"ability1":"blaze"}`}
	// res, err := db.NewInsert().Model(pokemon).Exec(ctx)
	// fmt.Print(res)

	// Delete
	// pokemons := []*Pokemon{{ID: 1}}
	// res, err := db.NewDelete().Model(&pokemons).WherePK().Exec(ctx)
	// fmt.Print(res)

	// Update
	// pokemon := &Pokemon{ID: 2, Description: "Hot Boi."}
	// res, err := db.NewUpdate().Model(pokemon).OmitZero().WherePK().Exec(ctx)
	// fmt.Print(res)

	// Find By ID
	// pokemon1 := new(Pokemon)
	// if err := db.NewSelect().Model(pokemon1).Where("id = ?", 2).Scan(ctx); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf(pokemon1.Name)

	// Find By Name
	// pokemon1 := new(Pokemon)
	// if err := db.NewSelect().Model(pokemon1).Where("name = ?", "Charmander").Scan(ctx); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf(pokemon1.Name)
}

type Pokemon struct {
	bun.BaseModel `bun:"table:Pokedex"`

	ID          int64 `bun:",pk,autoincrement"`
	Name        string
	Description string
	Category    string
	Type        string
	Abilities   string
}
