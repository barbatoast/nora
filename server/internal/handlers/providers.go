package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type ProviderHandler struct { DB *sqlx.DB }

type Provider struct {
	ID int64 `db:"id" json:"id"`
	FirmID *int64 `db:"firm_id" json:"firmId,omitempty"`
	Name string `db:"name" json:"name"`
	Model string `db:"model" json:"model"`
	BaseURL *string `db:"base_url" json:"baseUrl,omitempty"`
}

func (h ProviderHandler) List(w http.ResponseWriter, r *http.Request) {
	var rows []Provider
	h.DB.Select(&rows, `SELECT id, firm_id, name, model, base_url FROM providers ORDER BY id`)
	_ = json.NewEncoder(w).Encode(rows)
}

func (h ProviderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in Provider
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil { http.Error(w, err.Error(), 400); return }
	var id int64
	h.DB.QueryRow(`INSERT INTO providers (firm_id,name,model,base_url) VALUES ($1,$2,$3,$4) RETURNING id`, in.FirmID, in.Name, in.Model, in.BaseURL).Scan(&id)
	in.ID = id
	_ = json.NewEncoder(w).Encode(in)
}
