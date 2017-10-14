package model

import (
	"net/http"
	"errors"
	"time"

	"../global"
)

func CheckIsAdmin(w http.ResponseWriter, r *http.Request) error{
	if cookie, err := r.Cookie("auth"); err == nil {
		if cookie.Value == global.Auth.ShaPass {
			return nil
		}
	}
	return errors.New("Not Admin!")
}

func CheckAuth(w http.ResponseWriter, r *http.Request) error{
	r.ParseForm()
	login := r.Form.Get("login")
	password := r.Form.Get("password")
	if login == "" || password == ""{
		return errors.New("login or password is empty")
	}
	if login == global.Auth.Login && password == global.Auth.Password{
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "auth", Value: global.Auth.ShaPass, Expires: expiration}
		http.SetCookie(w, &cookie)
		return nil
	} else{
		return errors.New("Password incorrect")
	}
	return nil
}

func Logout(w http.ResponseWriter, r *http.Request) error{
	cookie := http.Cookie{Name: "auth", Value: ""}
	http.SetCookie(w, &cookie)
	return nil
}