package middleware

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// HttprouterWrapper  implementation is instructed by Nicolas Mérouze at https://www.nicolasmerouze.com/guide-routers-golang
func HttprouterWrapper(handler http.Handler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		context.WithValue(request.Context(), "params", params)
		handler.ServeHTTP(writer, request)
	}
}
