package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Acesso() *sql.DB {
	var conexao = "user=dba dbname=dba host=localhost password=123 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}
	return db
}
