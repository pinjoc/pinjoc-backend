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
    b.name AS base_token_name, 
    b.symbol AS base_token_symbol, 
    b.address AS base_token_address,
    CONCAT(MIN(t.price), ' ~ ', MAX(t.price)) AS price_range,
    CONCAT(MIN(m.month_name), ' ', MIN(m.year), ' ~ ', MAX(m.month_name), ' ', MAX(m.year)) AS maturity_range,
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
