// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    email
) VALUES (
    $1, $2, $3
) RETURNING user_id, username, password, email
`

type CreateUserParams struct {
	Username sql.NullString `json:"username"`
	Password sql.NullString `json:"password"`
	Email    sql.NullString `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password, arg.Email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Email,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username sql.NullString) error {
	_, err := q.db.ExecContext(ctx, deleteUser, username)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT user_id, username, password, email FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Password,
			&i.Email,
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

const getUserDetails = `-- name: GetUserDetails :one
SELECT user_id, username, password, email FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserDetails(ctx context.Context, username sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserDetails, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Email,
	)
	return i, err
}