package middleware

import (
	"context"
	"github.com/yun/UserManger/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/login", "/register":
			next.ServeHTTP(w, r)
			return
		}
		if cookie, err := r.Cookie("jwt_token"); err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else if claims, err := utils.ParseToken(cookie.Value); err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else {
			ctx := context.WithValue(r.Context(), "userID", claims.UserID)
			ctx = context.WithValue(ctx, "role", claims.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
