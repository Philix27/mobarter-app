
-- ! #SECTION Ads

-- name: CreateAd :one
INSERT INTO ads (
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

-- name: ListAds :many
SELECT * FROM ads
WHERE ad_status = $1 LIMIT 20;

-- name: GetAd :one
SELECT * FROM ads
WHERE id = $1 LIMIT 1;

-- name: UpdateAd :exec
UPDATE ads
  SET currency_pair = $2,
  limit_upper = $3,
  limit_lower = $4,
  rate = $5,
  instructions = $6
WHERE id = $1;

-- name: DeleteAd :exec
DELETE FROM ads
WHERE id = $1;