package models

type Pokemon struct {
	Name      string
	Abilities string
}

type PokemonDto struct {
	Name      string `json:"name" binding:"required"`
	Abilities string `json:"abilities" binding:"required"`
}

func (dto *PokemonDto) ToDto() *Pokemon {
	return &Pokemon{
		Name:      dto.Name,
		Abilities: dto.Abilities,
	}
}
