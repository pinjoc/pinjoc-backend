-- name: GetToken :many
SELECT price, amount, order_type
FROM tokenized t
JOIN token q ON t.quote_token_id = q.id
JOIN token b ON t.base_token_id = b.id
JOIN maturities m ON t.maturity_id = m.id
WHERE q.address = $1 AND b.address = $2 AND m.month = $3 AND m.year = $4 AND t.rate = $5;

-- name: GetAllToken :many
SELECT 
    q.name AS quote_token_name, 
    q.symbol AS quote_token_symbol, 
    q.address AS quote_token_address,
    q.icon AS quote_token_icon,
    b.name AS base_token_name, 
    b.symbol AS base_token_symbol, 
    b.address AS base_token_address,
    b.icon AS base_token_icon,
    CONCAT(MIN(t.price), ' ~ ', MAX(t.price)) AS price_range,
    CONCAT(
        (SELECT month_name FROM maturities WHERE id = MIN(m.id)), ' ', 
        (SELECT year FROM maturities WHERE id = MIN(m.id)), 
        ' ~ ', 
        (SELECT month_name FROM maturities WHERE id = MAX(m.id)), ' ', 
        (SELECT year FROM maturities WHERE id = MAX(m.id))
    ) AS maturity_range,
    SUM(t.volume) AS volume24h
FROM tokenized t
JOIN token q ON t.quote_token_id = q.id
JOIN token b ON t.base_token_id = b.id
JOIN maturities m ON t.maturity_id = m.id
GROUP BY q.id, b.id;

-- name: GetBasePrice :one
SELECT PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY price) AS best_price
FROM tokenized t
JOIN token q ON t.quote_token_id = q.id
JOIN token b ON t.base_token_id = b.id
JOIN maturities m ON t.maturity_id = m.id
WHERE q.address = $1 AND b.address = $2 AND m.month = $3 AND m.year = $4 AND t.rate = $5;

-- name: TokenAmount :exec
UPDATE tokenized
SET amount = amount + $1
WHERE id = $2;

-- name: TokenVolume :exec
UPDATE tokenized
SET volume = volume + $1
WHERE id = $2;

-- name: GetRandomToken :one
SELECT id FROM tokenized
ORDER BY RANDOM();

-- name: UpdateAmount :one
-- name: UpdateAvailable :one
UPDATE tokenized t
SET amount = $1
FROM token q, token b, maturities m
WHERE t.quote_token_id = q.id
AND t.base_token_id = b.id
AND t.maturity_id = m.id
AND q.address = $2
AND b.address = $3
AND m.month = $4
AND m.year = $5
AND t.order_type = $6
AND t.rate = $7
RETURNING t.id;