-- name: CreateUser :one
INSERT INTO users (
  name, email, role, password
) VALUES (
  $1, $2, $3, $4
)RETURNING id, created_at, *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
  SET name = $2,
      email = $3,
      role = $4,
      password = $5
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
