-- ! #SECTION bank_account
-- name: ListBankAccounts :many
SELECT * FROM bank_account
WHERE user_id = $1
ORDER BY bank_name;


-- name: GetBankAccount :one
SELECT * FROM bank_account
WHERE id = $1 LIMIT 1;
