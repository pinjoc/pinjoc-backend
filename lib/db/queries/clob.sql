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
    d.icon AS debt_token_icon,
    c.name AS collateral_token_name, 
    c.symbol AS collateral_token_symbol, 
    c.address AS collateral_address,
    c.icon AS collateral_token_icon,
    CONCAT(
        (SELECT month_name FROM maturities WHERE id = MIN(m.id)), ' ', 
        (SELECT year FROM maturities WHERE id = MIN(m.id)), 
        ' ~ ', 
        (SELECT month_name FROM maturities WHERE id = MAX(m.id)), ' ', 
        (SELECT year FROM maturities WHERE id = MAX(m.id))
    ) AS maturity_range,
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

-- name: UpdateAvailable :one
UPDATE orders o
SET available_token = $1
FROM token c, token d, maturities m
WHERE o.collateral_token_id = c.id
AND o.debt_token_id = d.id
AND o.maturity_id = m.id
AND c.address = $2
AND d.address = $3
AND m.month = $4
AND m.year = $5
AND o.order_type = $6
AND o.rate = $7
RETURNING o.id;


-- name: GetMaturitiesAndBestRate :many
SELECT 
    CONCAT(m.month_name, ' ', m.year) AS maturity,
    PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY o.rate) AS best_rate
FROM orders o
JOIN maturities m ON o.maturity_id = m.id
JOIN token c ON o.collateral_token_id = c.id
JOIN token d ON o.debt_token_id = d.id
WHERE c.address = $1 
AND d.address = $2
GROUP BY m.id, m.month_name, m.year
ORDER BY m.year ASC, m.month ASC;