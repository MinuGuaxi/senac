package modelo

import (
	"database/sql"
	"fmt"
	"log"
	dbconn "modulo/db"
)

type Usuarios struct {
	IDUsuarios int
	Nome       string
	Email      string
	Senha      string
}

func Inserir(nome string, email string, senha string) {
	db := dbconn.Acesso()

	stmt, err := db.Prepare("INSERT INTO usuarios (nome, email, senha) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println("Erro ao tentar salvar:", err)
		return
	}

	_, execErr := stmt.Exec(nome, email, senha)
	if execErr != nil {
		log.Println("Erro ao executar SQL:", execErr)
	} else {
		log.Println("Dados inseridos com sucesso:", nome)
	}

	defer db.Close()
}

func AutenticarUsuario(nome string, email string, senha string) (bool, error) {
	db := dbconn.Acesso()

	var senhaHash string
	query := "SELECT senha FROM usuarios WHERE email=$1"

	err := db.QueryRow(query, email).Scan(&senhaHash)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuário não encontrado:", email)
			return false, nil
		}
		fmt.Println("Erro ao executar Query:", err)
		return false, err
	}

	fmt.Println("Senha encontrada:", senhaHash)

	if senhaHash == senha {
		return true, nil
	}

	fmt.Println("Senha incorreta:", email)
	return false, nil
}
