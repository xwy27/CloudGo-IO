package router

import (
	"net/http"

	"github.com/unrolled/render"
)

type info struct {
	Author  string `json:"author"`
	Contact string `json:"contact"`
}

// InfoHandler returns the author info for info request
func InfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, info{
			Contact: "xuwy27@mail2.sysu.edu.cn",
			Author:  "xwy27"})
	}
}
