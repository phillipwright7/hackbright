-- name: CreateSale :one
INSERT INTO sales (
    car_id,
    owner_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSaleDetails :one
SELECT * FROM sales
WHERE sale_id = $1;

-- name: GetAllSales :many
SELECT * FROM sales;

-- name: DeleteSale :exec
DELETE FROM sales
WHERE sale_id = $1;

-- name: UpdateSale :exec
UPDATE sales 
SET car_id = $2, owner_id = $3
WHERE sale_id = $1;