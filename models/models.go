package models

import (
	"strconv"

	"github.com/deadman360/projetoLPSol/db"
)

type Corretor struct {
	Id, Creci int
	Nome      string
}

func Search(id string) Corretor {
	corr := Corretor{}
	dbs := db.DbConnect()
	defer dbs.Close()
	idn, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	query, err := dbs.Query("select * from solen where id=$1", idn)
	if err != nil {
		panic(err.Error())
	}
	var creci int
	var nome string
	for query.Next(){
	query.Scan(&id, &creci, &nome)

	corr.Nome = nome

	corr.Id = idn

	corr.Creci = creci
	}

	return corr
}
