-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :exec
INSERT INTO users (name, lastname, email, telephone, country) VALUES (?, ?, ?, ?, ?);
