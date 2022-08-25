package graph

import (
	"context"
	"strconv"

	"github.com/Gromitmugs/pokedex-graphql-sqlite/graph/db_model"
	"github.com/uptrace/bun"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Database struct {
	SqliteDB *bun.DB
}

type Resolver struct {
	DB Database
}

func (db *Database) FindById(id string, ctx context.Context) (*db_model.Pokemon, error) {
	pokemon1 := new(db_model.Pokemon) // return as a pointer
	if err := db.SqliteDB.NewSelect().Model(pokemon1).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return pokemon1, nil
}

func (db *Database) FindByName(name string, ctx context.Context) (*db_model.Pokemon, error) {
	pokemon1 := new(db_model.Pokemon) // return as a pointer
	if err := db.SqliteDB.NewSelect().Model(pokemon1).Where("name = ?", name).Scan(ctx); err != nil {
		return nil, err
	}
	return pokemon1, nil
}

func (db *Database) FindAllPokemon(ctx context.Context) ([]*db_model.Pokemon, error) {
	pokedex := make([]db_model.Pokemon, 0)

	if err := db.SqliteDB.NewSelect().Model(&pokedex).OrderExpr("id ASC").Scan(ctx); err != nil {
		return nil, err
	}
	allPokemon := []*db_model.Pokemon{}
	for _, element := range pokedex {
		newPokemon := element
		allPokemon = append(allPokemon, &newPokemon)
	}
	return allPokemon, nil
}

func (db *Database) DeletePokemon(ID string, ctx context.Context) error {
	id2, err3 := strconv.Atoi(ID)
	if err3 != nil {
		return err3
	}

	pokemon := []*db_model.Pokemon{{ID: int64(id2)}}

	_, err := db.SqliteDB.NewDelete().Model(&pokemon).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) AddPokemon(input *db_model.Pokemon, ctx context.Context) error {
	_, err := db.SqliteDB.NewInsert().Model(input).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdatePokemon(input *db_model.Pokemon, ctx context.Context) error {
	_, err := db.SqliteDB.NewUpdate().Model(input).OmitZero().WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
