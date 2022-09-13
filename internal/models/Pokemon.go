package models

type Pokemon struct {
	Id        uint64
	Name      string
	Abilities []string
	Type      string
}

type Pokemons struct {
	Pokemons []Pokemon
}
