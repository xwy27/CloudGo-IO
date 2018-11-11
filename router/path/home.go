package router

import (
	"net/http"

	"github.com/unrolled/render"
)

// HomeHandler renders the default home page to user
func HomeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "home", struct{}{})
	}
}
