package main

import (
	"net/http"

	pg "golang_practice/pages"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", pg.HomePage)
	http.ListenAndServe(":8000", nil)
}
