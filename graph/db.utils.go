package graph

import (
	"strconv"

	dbmodel "github.com/Gromitmugs/pokedex-graphql-sqlite/graph/db_model"
	gqlmodel "github.com/Gromitmugs/pokedex-graphql-sqlite/graph/gqlmodel"
)

func Gqlmodel2DBmodel(input gqlmodel.PokemonUpdateInput) (*dbmodel.Pokemon, error) {
	pokemon := &dbmodel.Pokemon{}

	id, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, err
	}
	pokemon.ID = int64(id)

	if input.Name != nil {
		pokemon.Name = *input.Name
	}
	if input.Description != nil {
		pokemon.Description = *input.Description
	}
	if input.Category != nil {
		pokemon.Category = *input.Category
	}
	if input.Type != nil {
		types := []string{}
		for _, e := range input.Type {
			types = append(types, *e)
		}
		pokemon.Type = types
	}
	if input.Abilities != nil {
		abilities := []string{}
		for _, e := range input.Abilities {
			abilities = append(abilities, *e)
		}
		pokemon.Abilities = abilities
	}

	return pokemon, nil
}
