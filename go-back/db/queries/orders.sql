-- name: ListOrders :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY created_at DESC;


-- name: CreateOrder :one
INSERT INTO orders (
  user_id,
  agent_id,
  currency_pair,
  amount,
  rate,
  order_type
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: UpdateOrder :exec
UPDATE orders
  SET currency_pair = $2,
  currency_pair = $3,
  amount = $4,
  rate = $5,
  rate = $6,
  order_type = $7,
  order_status = $8
WHERE id = $1;