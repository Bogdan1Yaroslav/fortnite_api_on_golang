package pages

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("templates/index.html")

	t.Execute(w, nil)
	// t.ExecuteTemplate(w, "index", nil)

}
