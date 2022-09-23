package adapter

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/lunetco/pokedex_api/config"
)

func New(conf config.Database) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.Username, conf.Password, conf.DBName)
	return gorm.Open("postgres", connStr)
}
