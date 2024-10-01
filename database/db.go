package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func DbConnect() {
	conexao := "host=localhost user=deadman360 dbname=corretores password=110300vv port=5433 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(conexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao iniciar Database")
	}

}
