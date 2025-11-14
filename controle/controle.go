package controle

import (
	"text/template"
	"net/http"
	modelo "modulo/modelo"
	"strconv"
)

var front *template.Template

func init(){
	var err error
	front, err = template.ParseGlob("template/*.html")
	if err != nil {
		panic("Erro ao carregar templates: " + err.Error())
	}
}

func Tela_Login(w http.ResponseWriter, r *http.Request){
	front.ExecuteTemplate(w, "Tela_Login", nil)
}

func Cadastro(w http.ResponseWriter, r *http.Request) {
	front.ExecuteTemplate(w, "Cadastro", nil)
}
func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		email := r.FormValue("email")
		senha := r.FormValue("senha")

		convertsenha, err := strconv.Atoi(senha)
		if err != nil {
			panic(err.Error())
		}

		modelo.Inserir(
			nome,
			email,
			convertsenha,
		)
	}

	http.Redirect(w, r, "/index.html", 301)

}
