package router

import (
	"github.com/BrownieBrown/bloggy/internal/api/handler"
	"net/http"
)

type Router struct {
	*http.ServeMux
}

type Middleware func(http.Handler) http.Handler

func NewRouter() *Router {
	return &Router{
		http.NewServeMux(),
	}
}

func (r *Router) Init(hh *handler.HealthHandler, eh *handler.ErrorHandler) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("GET /v1/readiness", hh.HealthCheck)
	r.HandleFunc("GET /v1/err", eh.HandleError)
}
