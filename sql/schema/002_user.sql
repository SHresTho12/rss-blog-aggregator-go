-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
ALTER TABLE users ADD COLUMN api_key VARCHAR(255) NOT NULL UNIQUE DEFAULT uuid_generate_v4();
-- +goose Down
ALTER TABLE users DROP COLUMN api_key;
