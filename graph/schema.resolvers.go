package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	dbmodel "github.com/Gromitmugs/pokedex-graphql-sqlite/graph/db_model"
	"github.com/Gromitmugs/pokedex-graphql-sqlite/graph/generated"
	gqlmodel "github.com/Gromitmugs/pokedex-graphql-sqlite/graph/gqlmodel"
)

// PokemonCreate is the resolver for the pokemonCreate field.
func (r *mutationResolver) PokemonCreate(ctx context.Context, input gqlmodel.PokemonCreateInput) (*dbmodel.Pokemon, error) {
	if input.ID != nil {
		return nil, fmt.Errorf("id must be null")
	}

	if _, err := r.DB.FindPokemonByName(input.Name, ctx); err == nil {
		return nil, fmt.Errorf("pokemon name already existed")
	}

	newPokemon := dbmodel.Pokemon{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Type:        input.Type,
		Abilities:   input.Abilities,
	}

	if err := r.DB.AddPokemon(&newPokemon, ctx); err != nil {
		return nil, err
	}

	return &newPokemon, nil
}

// PokemonUpdate is the resolver for the pokemonUpdate field.
func (r *mutationResolver) PokemonUpdate(ctx context.Context, input gqlmodel.PokemonUpdateInput) (*dbmodel.Pokemon, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("id must not be null")
	}

	if _, err := r.DB.FindPokemonById(input.ID, ctx); err != nil {
		return nil, err
	}

	pokemon, err := Gqlmodel2DBmodel(input)
	if err != nil {
		return nil, err
	}

	if err = r.DB.UpdatePokemon(pokemon, ctx); err != nil {
		return nil, err
	}

	return r.DB.FindPokemonById(input.ID, ctx)
}

// PokemonDelete is the resolver for the pokemonDelete field.
func (r *mutationResolver) PokemonDelete(ctx context.Context, id string) (bool, error) {
	if _, err := r.DB.FindPokemonById(id, ctx); err != nil {
		return false, err
	}

	if err := r.DB.DeletePokemon(id, ctx); err != nil {
		return false, err
	}

	return true, nil
}

// Pokedex is the resolver for the pokedex field.
func (r *queryResolver) Pokedex(ctx context.Context) ([]*dbmodel.Pokemon, error) {
	pokedex, err := r.DB.FindAllPokemon(ctx)
	if err != nil {
		return nil, err
	}

	return pokedex, nil
}

// PokemonByID is the resolver for the pokemonByID field.
func (r *queryResolver) PokemonByID(ctx context.Context, id string) (*dbmodel.Pokemon, error) {
	return r.DB.FindPokemonById(id, ctx)
}

// PokemonByName is the resolver for the pokemonByName field.
func (r *queryResolver) PokemonByName(ctx context.Context, name string) (*dbmodel.Pokemon, error) {
	return r.DB.FindPokemonByName(name, ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
