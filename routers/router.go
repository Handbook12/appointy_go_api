package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Routers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	s := r.PathPrefix("/api").Subrouter()
	AddPlacesRouter(s)

	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
