
-- ! #SECTION Ads

-- name: Advert_Create :one
INSERT INTO adverts (
  rate,  
  limit_lower, 
  limit_upper,
  currency_pair,
  instructions 
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: Advert_List :many
SELECT * FROM adverts
WHERE ad_status = $1 LIMIT 20;

-- name: GetAdvert :one
SELECT * FROM adverts
WHERE id = $1 LIMIT 1;

-- name: Advert_Update :exec
UPDATE adverts
  SET currency_pair = $2,
  limit_upper = $3,
  limit_lower = $4,
  rate = $5,
  instructions = $6
WHERE id = $1;

-- name: Advert_Delete :exec
DELETE FROM adverts
WHERE id = $1;