package main

import (
	"fmt"
	"github.com/lunetco/pokedex_api/config"
	"github.com/lunetco/pokedex_api/internal/controllers"
	"github.com/lunetco/pokedex_api/internal/router"
	"github.com/lunetco/pokedex_api/internal/service"
	"log"
	"net/http"
	"time"
)

func main() {

	appConfig := config.NewConfig()
	pokemonService := service.NewPokemonService()
	pokemonController := controllers.NewPokemonController(pokemonService)
	appRouter := router.NewRouter(pokemonController)

	address := fmt.Sprintf(":%d", appConfig.Server.Port)
	log.Printf("Starting internal %v", address)

	s := &http.Server{
		Addr:           address,
		Handler:        appRouter,
		ReadTimeout:    appConfig.Server.TimeoutRead * time.Second,
		WriteTimeout:   appConfig.Server.TimeoutWrite * time.Second,
		IdleTimeout:    appConfig.Server.TimeoutIdle * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
