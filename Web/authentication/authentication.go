package authentication

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseFiles("public/userinfo.html"))

type data struct {
	Info string
}

func init() {
	http.HandleFunc("/", root)
	//http.HandleFunc()

}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	currentuser, err := user.CurrentOAuth(c)

	if err != nil {
		http.Error(w, "OAuth Authorization", http.StatusUnauthorized)
		return
	}

	data1 := datastore.NewKey(c, "NEW DATA", "INFORMATION", 0, nil)
	data2 := datastore.NewKey(c, "NEW DATA", "INFROMATION", 0, nil)
	d = new(data)

	err := datastore.Get(c, data1, d)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	temp.ExecuteTemplate(w, "public/userinfo.html", currentuser)
}
