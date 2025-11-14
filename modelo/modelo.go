package modelo

import (
	"database/sql"
	"fmt"
	"log"
	"modulo/db"
)

type usuarios struct {
	id_usuarios int
	nome string
	email string
	senha string
}

func Inserir(nome string, email string, senha string) {
	db := db.Acesso()
	
	Inseri, err := db.Prepare("INSERT INTO usuarios (nome, email, senha) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println("Erro ao tentar salvar", err)
		return
	}
	_, execErr := Inseri.Exec(nome, email, senha)
	if execErr != nil {
		log.Println("Erro ao executar SQL: ", execErr)
	} else {
		log.Println("Dados inseridos com sucesso em: ", nome)
	}

	defer db.Close()
}

func AutenticarUsuario(nome string, email string, senha string) (bool, error) {
	db != db.Acesso()

	var senhaHash string
	query := "SELECT senha FROM usuarios WHERE email=$1"
	err := db.QueryRow(query, email).Scan(&senhaHash)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuário não encontrado: ", email)
			return false, nil
		}
		fmt.Println("Erro ao executar Query: ", err)
		return false, err
	}
	fmt.Println("Senha encontrada: ", senhaHash)
	if senhaHash == senha {
		return true, nil
	}
	fmt.Println("Senha incorreta: ", email)
	return false, nil
}
