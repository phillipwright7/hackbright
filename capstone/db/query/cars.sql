-- name: CreateCar :one
INSERT INTO cars (
    make,
    model,
    price,
    transmission,
    trim_level,
    color
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetCarDetails :one
SELECT * FROM cars
WHERE car_id = $1;

-- name: GetAllCars :many
SELECT * FROM cars;

-- name: DeleteCar :exec
DELETE FROM cars 
WHERE car_id = $1;

-- name: UpdateCar :exec
UPDATE cars
SET make = $2, model = $3, price = $4, transmission = $5, trim_level = $6, color = $7
WHERE car_id = $1;