package middleware

import (
	"net/http"
	"strings"
	"vapi-dashboard/server/config"
	"vapi-dashboard/server/internal/auth"
)

func Auth(cfg config.Config, allowed ...string) func(http.Handler) http.Handler {
	allowedSet := map[string]struct{}{}
	for _, r := range allowed { allowedSet[r]=struct{}{} }
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			h := r.Header.Get("Authorization")
			if !strings.HasPrefix(h, "Bearer ") { http.Error(w, "missing token", 401); return }
			tok := strings.TrimPrefix(h, "Bearer ")
			claims, err := auth.Parse(cfg.JWTSecret, tok)
			if err != nil { http.Error(w, "bad token", 401); return }
			r = SetUser(r, claims)
			if len(allowedSet)>0 { if _, ok := allowedSet[claims.Role]; !ok { http.Error(w, "forbidden", 403); return } }
			next.ServeHTTP(w, r)
		})
	}
}
