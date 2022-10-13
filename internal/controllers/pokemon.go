package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lunetco/pokedex_api/internal/models"
	"net/http"
	"strconv"
)

type PokemonService interface {
	GetPokemon(ctx context.Context, id uint64) (*models.Pokemon, error)
	ListPokemon(ctx context.Context) *models.Pokemon
}

type PokemonController struct {
	PokemonService PokemonService
}

func NewPokemonController(pokemonService PokemonService) *PokemonController {
	return &PokemonController{
		PokemonService: pokemonService,
	}
}

func (pc *PokemonController) GetPokemon(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 64)
	if err != nil {
		NewError(ctx, http.StatusUnprocessableEntity, errors.New("invalid id"))
		return
	}
	pokemon, err := pc.PokemonService.GetPokemon(ctx, id)
	if err != nil {
		NewError(ctx, http.StatusNotFound, errors.New("could not find Pokemon"))
		return
	}
	ctx.JSON(http.StatusOK, pokemon)
	return
}

func (pc *PokemonController) ListPokemon(ctx *gin.Context) {
	pokemons := pc.PokemonService.ListPokemon(ctx)
	ctx.JSON(http.StatusOK, pokemons)
}
