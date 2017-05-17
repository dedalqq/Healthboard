package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

const sessionKey = "sessionID"

type handler struct {
	param    *parameters
	tpl      *template.Template
	sl       *ServiceList
	sessions map[string]bool
}

var server http.Server

func initWebServer(sl *ServiceList, param *parameters) {

	t, err := template.ParseFiles("template.html")
	if nil != err {
		println(err.Error())
		exit()
		return
	}

	server = http.Server{
		Addr: fmt.Sprintf(":%d", param.webServerPort),
		Handler: handler{
			param:    param,
			tpl:      t,
			sl:       sl,
			sessions: map[string]bool{},
		},
	}

	go server.ListenAndServe()
}

// ServeHTTP -
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "/" != r.URL.Path {
		return
	}

	if h.httpLogin(w, r) {
		h.tpl.Execute(w, h.sl)
		return
	}

	h.tpl.Execute(w, nil)
}

func getSessionID() string {
	return fmt.Sprintf("%X", rand.Int63())
}

func (h handler) httpLogin(w http.ResponseWriter, r *http.Request) bool {
	cook, _ := r.Cookie(sessionKey)
	if nil == cook {
		cook = &http.Cookie{
			Name:  sessionKey,
			Value: getSessionID(),
		}

		http.SetCookie(w, cook)
		return false
	}

	if r.Method == "GET" {
		return h.sessions[cook.Value]
	}

	r.ParseForm()
	if r.PostForm.Get("login") == h.param.webLogin && r.PostForm.Get("password") == h.param.webPassword {
		h.sessions[cook.Value] = true
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return false
	}

	if r.PostForm.Get("logout") == "true" {
		h.sessions[cook.Value] = false
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return false
	}

	return false
}
