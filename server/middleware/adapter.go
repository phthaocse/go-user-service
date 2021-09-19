package middleware

import "net/http"

// Adapter Follow the instruction of Mat Ryer at https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
type Adapter func(handler http.Handler) http.Handler

func Adapt(handler http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		handler = adapter(handler)
	}
	return handler
}
