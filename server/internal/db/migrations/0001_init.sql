CREATE TABLE IF NOT EXISTS firms (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  slug TEXT NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  role TEXT NOT NULL CHECK (role IN ('admin','manager','agent')),
  firm_id BIGINT REFERENCES firms(id) ON DELETE SET NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS providers (
  id BIGSERIAL PRIMARY KEY,
  firm_id BIGINT REFERENCES firms(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  model TEXT NOT NULL,
  base_url TEXT,
  api_key TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS assistants (
  id BIGSERIAL PRIMARY KEY,
  firm_id BIGINT NOT NULL REFERENCES firms(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  provider_id BIGINT NOT NULL REFERENCES providers(id) ON DELETE RESTRICT,
  first_message TEXT NOT NULL DEFAULT '',
  system_prompt TEXT NOT NULL DEFAULT '',
  temperature REAL NOT NULL DEFAULT 0.7,
  max_tokens INT NOT NULL DEFAULT 250,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS phone_numbers (
  id BIGSERIAL PRIMARY KEY,
  firm_id BIGINT NOT NULL REFERENCES firms(id) ON DELETE CASCADE,
  e164 TEXT NOT NULL UNIQUE,
  label TEXT NOT NULL,
  assistant_id BIGINT REFERENCES assistants(id) ON DELETE SET NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_assistants_firm ON assistants(firm_id);
CREATE INDEX IF NOT EXISTS idx_providers_firm ON providers(firm_id);
CREATE INDEX IF NOT EXISTS idx_users_firm ON users(firm_id);
