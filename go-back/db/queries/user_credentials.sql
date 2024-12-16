-- name: Kyc_UpdateNin :exec
UPDATE kyc_credentials
  SET nin = $2
WHERE user_id = $1;


-- name: Kyc_UpdateBvn :exec
UPDATE kyc_credentials
  SET bvn = $2
WHERE user_id = $1;

-- name: Kyc_UpdateAddressAndUtilityBill :exec
UPDATE kyc_credentials
  SET house_address = $2,
  utility_bill = $3
WHERE user_id = $1;

