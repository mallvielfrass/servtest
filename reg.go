package main

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var cookieStore = sessions.NewCookieStore([]byte("secret"))

const cookieName = "MyCookie"

type sesKey int

const (
	sesKeyLogin sesKey = iota
)

func index(w http.ResponseWriter, r *http.Request) {
	ses, err := cookieStore.Get(r, cookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	login, ok := ses.Values[sesKeyLogin].(string)
	if !ok {
		login = "anonymous"
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintf(w, "you are "+login+"<br>"+`<a href="/login">get login</a>`)
}

func login(w http.ResponseWriter, r *http.Request) {
	ses, err := cookieStore.Get(r, cookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Seslog := ses.Values[sesKeyLogin].(string)
	if Seslog != "" {
		fmt.Println(Seslog)

	} else {

		ses.Values[sesKeyLogin] = StringGen(128)
		err = cookieStore.Save(r, w, ses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		Seslog = ses.Values[sesKeyLogin].(string)
	}
	w.Write([]byte("you are  " + Seslog))
}
func log() {

}
func main() {
	gob.Register(sesKey(0))

	router := mux.NewRouter()
	router.HandleFunc("/login", login)
	router.HandleFunc("/", index)
	http.ListenAndServe(":3000", router)
}
