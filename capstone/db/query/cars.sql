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