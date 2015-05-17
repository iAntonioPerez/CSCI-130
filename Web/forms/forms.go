package forms

import (
	"html/template"
	"net/http"
)

type user struct {
	Name  string
	Email string
}

var templates *template.Template

func init() {
	http.HandleFunc("/", login)
	http.HandleFunc("/info", info)
	templates = template.Must(template.ParseFiles("login.html", "info.html"))
}

func login(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	currentuser := user{r.FormValue("name"), r.FormValue("email")}
	//INFORMATION BEING PASSED: currentuser
	templates.ExecuteTemplate(w, "info.html", currentuser)
}
