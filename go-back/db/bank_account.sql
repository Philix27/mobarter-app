-- ! #SECTION bank_account
-- name: ListBankAccounts :many
SELECT * FROM bank_account
WHERE user_id = $1
ORDER BY bank_name;


-- name: GetBankAccount :one
SELECT * FROM bank_account
WHERE id = $1 LIMIT 1;



-- name: CreateBankAccount :one
INSERT INTO bank_account (
  user_id,
  bank_name,  
  account_name, 
  account_no
) VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: DeleteBankAccount :exec
DELETE FROM bank_account
WHERE id = $1;