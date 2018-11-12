package router

import (
	"net/http"

	"github.com/unrolled/render"
)

// Index info struct
type indexInfo struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// LoginHandler returns the login page with GET method and
// deals with a login form and returns the result with POST method
func LoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			formatter.HTML(w, http.StatusOK, "login", struct{}{})
		} else if req.Method == "POST" {
			var email, password []string
			req.ParseForm()
			for k, v := range req.Form {
				switch k {
				case "email":
					email = v
				case "password":
					password = v
				}
			}
			formatter.HTML(w, http.StatusOK, "index", indexInfo{
				Email:    email[0],
				Password: password[0],
			})
		}
	}
}
