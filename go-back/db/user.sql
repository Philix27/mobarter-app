-- ! #SECTION user

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  email 
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  SET first_name = $2,
  last_name = $3,
  phone = $4
WHERE id = $1;

-- name: UpdateKyc_credentials :exec
UPDATE users
  SET first_name = $2,
  last_name = $3,
  phone = $4
WHERE id = $1;