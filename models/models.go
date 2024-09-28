package models

import (
	"strconv"

	"github.com/deadman360/projetoLPSol/db"
)

type Corretor struct {
	Id   int
	Nome string
}

func Search(id string) Corretor {
	corr := Corretor{}
	dbs := db.DbConnect()
	defer dbs.Close()
	query, err := dbs.Query("select * from solen where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	var nome string
	for query.Next() {
		query.Scan(&id, &nome)

		corr.Nome = nome

		corr.Id, err = strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}

	}

	return corr
}
