package controle

import (
	modelo "modulo/modelo"
	"net/http"
	"text/template"
)

var front *template.Template

func init() {
	var err error
	front, err = template.ParseGlob("template/*.html")
	if err != nil {
		panic("Erro ao carregar templates: " + err.Error())
	}
}

func Tela_Login(w http.ResponseWriter, r *http.Request) {
	front.ExecuteTemplate(w, "Tela_Login", nil)
}

func Cadastro(w http.ResponseWriter, r *http.Request) {
	front.ExecuteTemplate(w, "Cadastro", nil)
}

func Tela(w http.ResponseWriter, r *http.Request) {
	front.ExecuteTemplate(w, "Tela", nil)
}

func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		email := r.FormValue("email")
		senha := r.FormValue("senha")

		modelo.Inserir(
			nome,
			email,
			senha,
		)
	}

	http.Redirect(w, r, "/Tela", http.StatusSeeOther)

}
