-- name: CreateMovie :one
INSERT INTO movies (
    title,
    overview,
    release_date, 
    poster_url
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetMovieDetails :one
SELECT * FROM movies
WHERE title = $1;

-- name: GetAllMovies :many
SELECT * FROM movies;

-- name: DeleteMovie :exec
DELETE FROM movies WHERE title = $1;