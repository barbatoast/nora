package middleware

import (
	"context"
	"net/http"
	"vapi-dashboard/server/internal/auth"
)

type key string
const userKey key = "user-claims"

func SetUser(r *http.Request, c *auth.Claims) *http.Request { return r.WithContext(context.WithValue(r.Context(), userKey, c)) }
func GetUser(r *http.Request) *auth.Claims { v := r.Context().Value(userKey); if v==nil { return nil }; return v.(*auth.Claims) }
