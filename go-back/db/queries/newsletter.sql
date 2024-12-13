-- ! #SECTION Newsletter
-- name: ListNewsletter :many
SELECT * FROM newsletter
ORDER BY email;


-- name: CreateNewsletter :one
INSERT INTO newsletter (
  email
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteNewsletter :exec
DELETE FROM newsletter
WHERE email = $1;