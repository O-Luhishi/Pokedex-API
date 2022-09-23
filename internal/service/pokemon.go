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
		return &models.Pokemon{
			Name:      "Charmander",
			Abilities: "",
		}, nil
	} else if id == 2 {
		return &models.Pokemon{
			Name:      "Bulbasaur",
			Abilities: "abilities",
		}, nil
	} else {
		// We'll add error handling later when we search through a proper datasource!
		return nil, errors.New("could not find pokemon")
	}
}

func (ps *PokemonService) ListPokemon(ctx context.Context) *models.Pokemon {
	return nil
}
