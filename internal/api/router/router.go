package router

import "net/http"

type Router struct {
	*http.ServeMux
}

type Middleware func(http.Handler) http.Handler

func NewRouter() *Router {
	return &Router{
		http.NewServeMux(),
	}
}

func (r *Router) Init() {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
