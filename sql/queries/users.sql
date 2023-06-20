-- name: CreateUser :one
INSERT INTO users (id, name, email, password, created_at, updated_at,api_key) VALUES ($1, $2, $3, $4, $5, $6,uuid_generate_v4()) RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;
