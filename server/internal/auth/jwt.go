package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64  `json:"uid"`
	Role   string `json:"role"`
	FirmID *int64 `json:"firmId,omitempty"`
	jwt.RegisteredClaims
}

func Sign(secret string, userID int64, role string, firmID *int64, ttl time.Duration) (string, error) {
	claims := Claims{UserID:userID, Role:role, FirmID:firmID, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl))}}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

func Parse(secret, token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token)(interface{}, error){ return []byte(secret), nil })
	if err != nil { return nil, err }
	if c, ok := t.Claims.(*Claims); ok && t.Valid { return c, nil }
	return nil, jwt.ErrTokenInvalidClaims
}
