package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/.well-known/appspecific/com.chrome.devtools.json" {
			return
		}
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		//log.Printf("[%s] %s %s %s", r.Method, r.RequestURI, r.Proto, duration)
		log.Printf("[%s] %s %s %s", r.Method, r.URL.Path, r.RemoteAddr, duration)
	})
}
