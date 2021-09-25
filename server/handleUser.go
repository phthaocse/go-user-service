package server

import (
	"net/http"
)

func (srv *server) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := srv.decode(w, r, nil); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
