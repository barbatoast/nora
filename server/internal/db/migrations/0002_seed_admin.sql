INSERT INTO firms (name, slug)
VALUES ('Demo Law', 'demo-law')
ON CONFLICT DO NOTHING;

-- password: adminadmin
INSERT INTO users (email, password_hash, role, firm_id)
SELECT 'admin@demo.law', '$2a$12$k9GkXIrQ6bqz5uI2xXz8S.6iHt0m8b4C3H2nJzH0aU8Z8o0q2mW8S', 'admin', NULL
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email='admin@demo.law');
