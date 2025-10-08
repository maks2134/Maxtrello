CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

INSERT INTO users (email, password, name)
VALUES ('test@example.com', 'ifjosafjioefgj', 'Test User')
ON CONFLICT (email) DO NOTHING;