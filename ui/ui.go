package ui

import (
	"net"
	"net/http"
	"time"
	"strings"
	"crypto/sha256"
	"encoding/base64"

	"../global"
	"../controller"
	"../middleware"
)

func Start(listener net.Listener) {
	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
		}

	http.Handle("/", http.FileServer(http.Dir("./Public")))
	http.HandleFunc("/index", makeHandler(controller.Index, "GET", true))
	http.HandleFunc("/login", makeHandler(controller.Login, "POST, GET", false))
	http.HandleFunc("/logout", makeHandler(controller.Logout, "GET", true))

	hasher := sha256.New()
	hasher.Write([]byte(global.Auth.Login + global.Auth.Password))
	global.Auth.ShaPass = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	go server.Serve(listener)
}


func makeHandler(fn func(http.ResponseWriter, *http.Request), method string, secret bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(method, r.Method) < 0{
			http.NotFound(w, r)
			return
		}
		if secret == true{
			w, r = middleware.AuthMiddleware(w, r)
		}

		fn(w, r)
	}
}
