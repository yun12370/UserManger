package middleware

import (
	"context"
	"github.com/yun/UserManger/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		publicPaths := map[string]bool{
			"/login":    true,
			"/register": true,
		}

		if cookie, err := r.Cookie("jwt_token"); err == nil {
			if claims, err := utils.ParseToken(cookie.Value); err == nil {
				if publicPaths[path] {
					http.Redirect(w, r, "/index", http.StatusSeeOther)
					return
				}
				ctx := context.WithValue(r.Context(), "userID", claims.UserID)
				ctx = context.WithValue(ctx, "role", claims.Role)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		if publicPaths[path] {
			next.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})
}
