package db

import "database/sql"

func DbConnect() *sql.DB {
	conexao := "user=deadman360 dbname=corretores password=110300vv host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
