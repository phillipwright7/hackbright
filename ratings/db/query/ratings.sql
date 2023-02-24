-- name: CreateRating :one
INSERT INTO ratings (
    movie_id,
    user_id,
    score
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetRatingDetails :one
SELECT * FROM ratings
WHERE movie_id = $1 AND user_id = $2;

-- name: UpdateRating :one
UPDATE ratings
SET score = $1
WHERE movie_id = $2 AND user_id = $3
RETURNING *;

-- name: DeleteRating :exec
DELETE FROM ratings
WHERE movie_id = $1 AND user_id = $2;

-- name: GetMovieRatings :many
SELECT * FROM ratings
WHERE movie_id = $1;

-- name: GetUserRatings :many
SELECT score FROM ratings
WHERE user_id = $1;