package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type NumberHandler struct { DB *sqlx.DB }

type Number struct {
	ID int64 `db:"id" json:"id"`
	FirmID int64 `db:"firm_id" json:"firmId"`
	E164 string `db:"e164" json:"e164"`
	Label string `db:"label" json:"label"`
	AssistantID *int64 `db:"assistant_id" json:"assistantId,omitempty"`
}

func (h NumberHandler) List(w http.ResponseWriter, r *http.Request) {
	var rows []Number
	h.DB.Select(&rows, `SELECT id, firm_id, e164, label, assistant_id FROM phone_numbers ORDER BY id`)
	_ = json.NewEncoder(w).Encode(rows)
}

func (h NumberHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in Number
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil { http.Error(w, err.Error(), 400); return }
	var id int64
	h.DB.QueryRow(`INSERT INTO phone_numbers (firm_id,e164,label,assistant_id) VALUES ($1,$2,$3,$4) RETURNING id`, in.FirmID, in.E164, in.Label, in.AssistantID).Scan(&id)
	in.ID = id
	_ = json.NewEncoder(w).Encode(in)
}
