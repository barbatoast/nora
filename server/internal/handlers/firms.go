package handlers

import (
	"encoding/json"
	"net/http"
	"vapi-dashboard/server/internal/middleware"
	"github.com/jmoiron/sqlx"
)

type FirmHandler struct { DB *sqlx.DB }

type Firm struct{ ID int64 `db:"id" json:"id"`; Name string `db:"name" json:"name"`; Slug string `db:"slug" json:"slug"` }

func (h FirmHandler) List(w http.ResponseWriter, r *http.Request) {
	claims := middleware.GetUser(r)
	rows := []Firm{}
	if claims != nil && claims.Role == "admin" {
		h.DB.Select(&rows, `SELECT id,name,slug FROM firms ORDER BY id`)
	} else if claims != nil && claims.FirmID != nil {
		h.DB.Select(&rows, `SELECT id,name,slug FROM firms WHERE id=$1`, *claims.FirmID)
	}
	_ = json.NewEncoder(w).Encode(rows)
}
