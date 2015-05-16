package cookies

import (
	"encoding/json"
	//"fmt"
	"html/template"
	"net/http"
)

var temp *template.Template

type id struct {
	Number int
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/displayinfo", displayinfo)
	http.HandleFunc("/cookie", page)
	temp = template.Must(template.ParseFiles("info.html", "displyinfo.html", "cookie.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "info.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func displayinfo(w http.ResponseWriter, r *http.Request) {
	user := id{Number: r.FormValue("number")}

	data, _ := json.Marshal(user)

	cookie := http.Cookie{
		Name:  "userid",
		Value: string(data)}

	http.SetCookie(w, &cookie)
	err := temp.ExecuteTemplate(w, displayinfo.html, struct {
		Number int
	}{r.FormValue("number")})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func page(w http.ResponseWriter, r *http.Request) {
	var user id
	cookie, _ := r.Cookie("test")
	data := []byte(cookie.Value)
	err := jsonUnmarshal(data, &user)
	err = temp.ExecuteTemplate(w, "cookie.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
