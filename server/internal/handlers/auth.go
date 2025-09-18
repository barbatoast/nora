package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"vapi-dashboard/server/config"
	"vapi-dashboard/server/internal/auth"

	"github.com/jmoiron/sqlx"
)

type AuthHandler struct { DB *sqlx.DB; Cfg config.Config }

type loginReq struct{ Email, Password string }

type loginResp struct{ Token string `json:"token"` }

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var in loginReq
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil { http.Error(w, err.Error(), 400); return }
	var user struct{ ID int64 `db:"id"`; Hash string `db:"password_hash"`; Role string `db:"role"`; FirmID sql.NullInt64 `db:"firm_id"` }
	if err := h.DB.Get(&user, `SELECT id, password_hash, role, firm_id FROM users WHERE email=$1`, in.Email); err != nil { http.Error(w, "invalid credentials", 401); return }
	
	hash := strings.TrimSpace(user.Hash)
	if err := auth.CheckPassword(hash, in.Password); err != nil { http.Error(w, "invalid credentials", 401); return }
	var firmPtr *int64; if user.FirmID.Valid { v := user.FirmID.Int64; firmPtr = &v }
	tok, _ := auth.Sign(h.Cfg.JWTSecret, user.ID, user.Role, firmPtr, 24*time.Hour)
	_ = json.NewEncoder(w).Encode(loginResp{Token: tok})
}
