-- ! #SECTION airtime
-- name: ListAirtime :many
SELECT * FROM airtime
ORDER BY network;

-- ! #SECTION data_plans
-- name: ListDataPlans :many
SELECT * FROM data_plans
ORDER BY network;
