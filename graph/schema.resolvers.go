package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Gromitmugs/pokedex-graphql-sqlite/graph/db_model"
	"github.com/Gromitmugs/pokedex-graphql-sqlite/graph/generated"
	"github.com/Gromitmugs/pokedex-graphql-sqlite/graph/model"
)

// PokemonCreate is the resolver for the pokemonCreate field.
func (r *mutationResolver) PokemonCreate(ctx context.Context, input model.PokemonInput) (*db_model.Pokemon, error) {
	if input.ID != nil {
		return nil, fmt.Errorf("id must be null")
	}

	_, err_name := r.DB.FindByName(input.Name, ctx)
	if err_name == nil {
		return nil, fmt.Errorf("pokemon name already existed")
	}

	newPokemon := db_model.Pokemon{
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Type:        input.Type,
		Abilities:   input.Abilities,
	}

	err := r.DB.AddPokemon(&newPokemon, ctx)
	if err != nil {
		return nil, err
	}

	return &newPokemon, nil
}

// PokemonUpdate is the resolver for the pokemonUpdate field.
func (r *mutationResolver) PokemonUpdate(ctx context.Context, input model.PokemonInput) (*db_model.Pokemon, error) {
	if input.ID == nil {
		return nil, fmt.Errorf("id must not be null")
	}

	_, err2 := r.DB.FindById(*input.ID, ctx)
	if err2 != nil {
		return nil, err2
	}
	id2, err3 := strconv.Atoi(*input.ID)
	if err3 != nil {
		return nil, err3
	}

	pokemon := db_model.Pokemon{ID: int64(id2),
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Type:        input.Type,
		Abilities:   input.Abilities,
	}
	err := r.DB.UpdatePokemon(&pokemon, ctx)
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}

// PokemonDelete is the resolver for the pokemonDelete field.
func (r *mutationResolver) PokemonDelete(ctx context.Context, id string) (bool, error) {
	_, err2 := r.DB.FindById(id, ctx)

	if err2 != nil {
		return false, err2
	}

	err := r.DB.DeletePokemon(id, ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Pokedex is the resolver for the pokedex field.
func (r *queryResolver) Pokedex(ctx context.Context) ([]*db_model.Pokemon, error) {
	pokedex, err := r.DB.FindAllPokemon(ctx)
	if err != nil {
		return nil, err
	}

	return pokedex, nil
}

// PokemonByID is the resolver for the pokemonByID field.
func (r *queryResolver) PokemonByID(ctx context.Context, id string) (*db_model.Pokemon, error) {
	return r.DB.FindById(id, ctx)
}

// PokemonByName is the resolver for the pokemonByName field.
func (r *queryResolver) PokemonByName(ctx context.Context, name string) (*db_model.Pokemon, error) {
	return r.DB.FindByName(name, ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
