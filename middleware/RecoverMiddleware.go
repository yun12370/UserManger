package middleware

import (
	"encoding/json"
	"github.com/yun/UserManger/utils"
	"log"
	"net/http"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Recover] Panic: %v", err)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(utils.Fail[any](http.StatusInternalServerError, "系统错误"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
