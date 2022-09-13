package service

import (
	"context"
	"errors"
	"github.com/lunetco/pokedex_api/internal/models"
)

type PokemonService struct {
	// For now, it's empty
}

func NewPokemonService() *PokemonService {
	return &PokemonService{}
}

func (ps *PokemonService) GetPokemon(ctx context.Context, id uint64) (*models.Pokemon, error) {
	if id == 1 {
		abilities := []string{"Blaze", "Dodge"}
		return &models.Pokemon{
			Id:        0,
			Name:      "Charmander",
			Abilities: abilities,
			Type:      "Fire",
		}, nil
	} else if id == 2 {
		abilities := []string{"Overgrow", "Dodge"}
		return &models.Pokemon{
			Id:        2,
			Name:      "Bulbasaur",
			Abilities: abilities,
			Type:      "Grass",
		}, nil
	} else {
		// We'll add error handling later when we search through a proper datasource!
		return nil, errors.New("could not find pokemon")
	}
}

func (ps *PokemonService) ListPokemon(ctx context.Context) *models.Pokemons {
	return &models.Pokemons{Pokemons: []models.Pokemon{
		{
			Id:        0,
			Name:      "Charmander",
			Abilities: []string{"Blaze", "Dodge"},
			Type:      "Fire",
		},
		{
			Id:        2,
			Name:      "Bulbasaur",
			Abilities: []string{"Overgrow", "Dodge"},
			Type:      "Grass",
		},
	}}
}
