package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type AssistantHandler struct { DB *sqlx.DB }

type Assistant struct {
	ID int64 `db:"id" json:"id"`
	FirmID int64 `db:"firm_id" json:"firmId"`
	Name string `db:"name" json:"name"`
	ProviderID int64 `db:"provider_id" json:"providerId"`
	FirstMessage string `db:"first_message" json:"firstMessage"`
	SystemPrompt string `db:"system_prompt" json:"systemPrompt"`
	Temperature float32 `db:"temperature" json:"temperature"`
	MaxTokens int `db:"max_tokens" json:"maxTokens"`
}

func (h AssistantHandler) List(w http.ResponseWriter, r *http.Request) {
	var rows []Assistant
	h.DB.Select(&rows, `SELECT id, firm_id, name, provider_id, first_message, system_prompt, temperature, max_tokens FROM assistants ORDER BY id`)
	_ = json.NewEncoder(w).Encode(rows)
}

func (h AssistantHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in Assistant
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil { http.Error(w, err.Error(), 400); return }
	var id int64
	h.DB.QueryRow(`INSERT INTO assistants (firm_id,name,provider_id,first_message,system_prompt,temperature,max_tokens) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`, in.FirmID, in.Name, in.ProviderID, in.FirstMessage, in.SystemPrompt, in.Temperature, in.MaxTokens).Scan(&id)
	in.ID = id
	_ = json.NewEncoder(w).Encode(in)
}
