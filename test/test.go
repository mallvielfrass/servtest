package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := "login.html" //Страница входа
		if len(r.Header["Cookie"]) != 0 && r.Header["Cookie"][0] == "auth=your_MD5_cookies" {
			url = "index.html" //Страница после успешной авторизации
		}
		t, _ := template.ParseFiles(url)
		t.Execute(w, "")
	})
	http.Handle("/js/", http.FileServer(http.Dir("/")))
	http.ListenAndServe(":8000", nil)
}
