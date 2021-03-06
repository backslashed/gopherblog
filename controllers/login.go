package controller

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/gbbr/gopherblog/models"
	"net/http"
	"strings"
)

// Displays the login form template. Interprets both GET and POST
func Login(w http.ResponseWriter, r *http.Request) {
	returnTo := r.URL.Query().Get("return")
	if returnTo == "" {
		returnTo = "/manage"
	}

	tplData := struct{ Msg, ReturnUrl string }{
		ReturnUrl: returnTo,
	}

	if r.Method == "POST" {
		// this will redirect us if login succeeds
		validateLoginForm(w, r)
		// otherwise, display message
		tplData.Msg = "Invalid login and password combination."
	}

	tpl.ExecuteTemplate(w, "login", tplData)
}

// Validates the login form's POST and redirects the user if login
// is correct
func validateLoginForm(w http.ResponseWriter, r *http.Request) {
	md5pass := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))

	user := &models.User{
		Email:    r.FormValue("login"),
		Password: md5pass,
	}

	if user.LoginCorrect() {
		remoteIp := strings.Split(r.RemoteAddr, ":")[0]
		origin := []byte(string(user.Email) + remoteIp + r.UserAgent())
		val := fmt.Sprintf("%d:%x", user.Id, sha256.Sum256(origin))
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: val,
		})

		http.Redirect(w, r, r.FormValue("redirectUrl"), http.StatusFound)
	}
}
