package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lunetco/pokedex_api/internal/controllers"
	middlewares "github.com/lunetco/pokedex_api/internal/middleware"
)

func NewRouter(pokemonController *controllers.PokemonController) *gin.Engine {
	app := gin.Default()
	app.Use(gin.Recovery())
	app.NoRoute(middlewares.NoRouteHandler())

	pokemonRoutes := app.Group("api/v1")
	pokemonRoutes.Use()
	{
		pokemonRoutes.GET("pokemon/:id", pokemonController.GetPokemon)
		pokemonRoutes.GET("pokemon", pokemonController.ListPokemon)
	}
	return app
}
