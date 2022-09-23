package models

type Pokemon struct {
	Name      string `gorm:"column:pokemon_name;not null; default:null"`
	Abilities string `gorm:"column:abilities;not null; default:null"`
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
