package router

import (
	"net/http"
)

// DevelopHandler renders the not implemented hint
func DevelopHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
	}
}
