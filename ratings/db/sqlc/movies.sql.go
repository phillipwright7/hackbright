// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: movies.sql

package db

import (
	"context"
	"database/sql"
)

const createMovie = `-- name: CreateMovie :one
INSERT INTO movies (
    title,
    overview,
    release_date, 
    poster_url
) VALUES (
    $1, $2, $3, $4
) RETURNING movie_id, title, overview, release_date, poster_url
`

type CreateMovieParams struct {
	Title       sql.NullString `json:"title"`
	Overview    sql.NullString `json:"overview"`
	ReleaseDate sql.NullTime   `json:"release_date"`
	PosterUrl   sql.NullString `json:"poster_url"`
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, createMovie,
		arg.Title,
		arg.Overview,
		arg.ReleaseDate,
		arg.PosterUrl,
	)
	var i Movie
	err := row.Scan(
		&i.MovieID,
		&i.Title,
		&i.Overview,
		&i.ReleaseDate,
		&i.PosterUrl,
	)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE FROM movies WHERE title = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, title sql.NullString) error {
	_, err := q.db.ExecContext(ctx, deleteMovie, title)
	return err
}

const getAllMovies = `-- name: GetAllMovies :many
SELECT movie_id, title, overview, release_date, poster_url FROM movies
`

func (q *Queries) GetAllMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, getAllMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.MovieID,
			&i.Title,
			&i.Overview,
			&i.ReleaseDate,
			&i.PosterUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMovieDetails = `-- name: GetMovieDetails :one
SELECT movie_id, title, overview, release_date, poster_url FROM movies
WHERE title = $1
`

func (q *Queries) GetMovieDetails(ctx context.Context, title sql.NullString) (Movie, error) {
	row := q.db.QueryRowContext(ctx, getMovieDetails, title)
	var i Movie
	err := row.Scan(
		&i.MovieID,
		&i.Title,
		&i.Overview,
		&i.ReleaseDate,
		&i.PosterUrl,
	)
	return i, err
}
