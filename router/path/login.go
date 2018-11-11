package router

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"
)

// LoginHandler deals with a login form and renders a default login hint
func LoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// TODO: deal with login form
		fmt.Println("user: ")
		fmt.Println("password: ")
	}
}
