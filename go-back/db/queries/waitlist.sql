-- ! #SECTION waitlist
-- name: ListWaitlist :many
SELECT * FROM waitlist
ORDER BY email;


-- name: CreateWaitlist :one
INSERT INTO waitlist (
  email
) VALUES (
  $1
)
RETURNING *;
