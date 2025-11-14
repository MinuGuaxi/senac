package rotas

import (
	"modulo/controle"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controle.Tela_Login)
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	http.HandleFunc("/Cadastro", controle.Cadastro)
	http.HandleFunc("/Inserir", controle.Inserir)
}
