package routes

import (
	"fmt"
	"net/http"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		Mux: http.NewServeMux(),
	}
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.handleWithMethod(http.MethodPost, path, handler)
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.handleWithMethod(http.MethodGet, path, handler)
}

func (r *Router) handleWithMethod(method, path string, handler http.HandlerFunc) {
	r.Mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			http.Error(w, fmt.Sprintf("method %s not allowed", req.Method), http.StatusMethodNotAllowed)
			return
		}
		handler(w, req)
	})
}
