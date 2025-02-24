-- name: GetCLOB :many
SELECT rate, available_token, order_type
FROM orders o
JOIN token c ON o.collateral_token_id = c.id
JOIN token d ON o.debt_token_id = d.id
JOIN maturities m ON o.maturity_id = m.id
WHERE c.address = $1 AND d.address = $2 AND m.month = $3 AND m.year = $4;

-- name: GetAvailableToken :many
SELECT 
    d.name AS debt_token_name, 
    d.symbol AS debt_token_symbol, 
    d.address AS debt_token_address,
    c.name AS collateral_token_name, 
    c.symbol AS collateral_token_symbol, 
    c.address AS collateral_address,
    CONCAT(MIN(m.month_name), ' ', MIN(m.year), ' ~ ', MAX(m.month_name), ' ', MAX(m.year)) AS maturity_range,
    CONCAT(MIN(o.rate), '% ~ ', MAX(o.rate), '%') AS rate_range,
    SUM(CASE WHEN o.order_type = 'LEND' THEN o.available_token ELSE 0 END) AS lending_vault,
    SUM(CASE WHEN o.order_type = 'BORROW' THEN o.available_token ELSE 0 END) AS borrow_vault
FROM orders o
JOIN token c ON o.collateral_token_id = c.id
JOIN token d ON o.debt_token_id = d.id
JOIN maturities m ON o.maturity_id = m.id
GROUP BY d.id, c.id;


-- name: GetBestRate :one
SELECT PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY rate) AS best_rate
FROM orders o
JOIN token c ON o.collateral_token_id = c.id
JOIN token d ON o.debt_token_id = d.id
JOIN maturities m ON o.maturity_id = m.id
WHERE c.address = $1 AND d.address = $2 AND m.month = $3 AND m.year = $4;

-- name: UpdateTokenAvailable :exec
UPDATE orders
SET available_token = available_token + $1
WHERE id = $2;

-- name: GetRandomOrder :one
SELECT id FROM orders
ORDER BY RANDOM();