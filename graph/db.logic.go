package graph

import (
	"context"
	"strconv"

	dbmodel "github.com/Gromitmugs/pokedex-graphql-sqlite/graph/db_model"
	"github.com/uptrace/bun"
)

type Database struct {
	Postgres *bun.DB
}

func (db *Database) FindPokemonById(id string, ctx context.Context) (*dbmodel.Pokemon, error) {
	pokemon1 := new(dbmodel.Pokemon) // return as a pointer
	if err := db.Postgres.NewSelect().Model(pokemon1).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return pokemon1, nil
}

func (db *Database) FindPokemonByName(name string, ctx context.Context) (*dbmodel.Pokemon, error) {
	pokemon1 := new(dbmodel.Pokemon) // return as a pointer
	if err := db.Postgres.NewSelect().Model(pokemon1).Where("name = ?", name).Scan(ctx); err != nil {
		return nil, err
	}
	return pokemon1, nil
}

func (db *Database) FindAllPokemon(ctx context.Context) ([]*dbmodel.Pokemon, error) {
	pokedex := make([]dbmodel.Pokemon, 0)

	if err := db.Postgres.NewSelect().Model(&pokedex).OrderExpr("id ASC").Scan(ctx); err != nil {
		return nil, err
	}
	allPokemon := []*dbmodel.Pokemon{}
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

	pokemon := []*dbmodel.Pokemon{{ID: int64(id2)}}

	_, err := db.Postgres.NewDelete().Model(&pokemon).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) AddPokemon(input *dbmodel.Pokemon, ctx context.Context) error {
	_, err := db.Postgres.NewInsert().Model(input).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdatePokemon(input *dbmodel.Pokemon, ctx context.Context) error {
	_, err := db.Postgres.NewUpdate().Model(input).OmitZero().WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
