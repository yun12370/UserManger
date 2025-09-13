package middleware

import "net/http"

func Chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		count := len(middlewares) - 1
		for i := count; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
