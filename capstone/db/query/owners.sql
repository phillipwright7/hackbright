-- name: CreateOwner :one
INSERT INTO owners (
    first_name,
    last_name,
    phone_number,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetOwnerDetails :one
SELECT * FROM owners
WHERE owner_id = $1;

-- name: GetAllOwners :many
SELECT * FROM owners;

-- name: DeleteOwner :exec
DELETE FROM owners
WHERE owner_id = $1;

-- name: UpdateOwner :exec
UPDATE owners
SET first_name = $2, last_name = $3, phone_number = $4, email = $5
WHERE owner_id = $1;