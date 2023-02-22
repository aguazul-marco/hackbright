// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: movies.sql

package db

import (
	"context"
)

const createMovie = `-- name: CreateMovie :one
INSERT INTO movies (
  title, 
  overview,
  release_date,
  poster_url
) VALUES (
  $1, $2, $3, $4
) RETURNING id, title, overview, release_date, poster_url
`

type CreateMovieParams struct {
	Title       string `json:"title"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"release_date"`
	PosterUrl   string `json:"poster_url"`
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
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.ReleaseDate,
		&i.PosterUrl,
	)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE FROM movies
WHERE id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMovie, id)
	return err
}

const getMovie = `-- name: GetMovie :one
SELECT id, title, overview, release_date, poster_url FROM movies
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMovie(ctx context.Context, id int64) (Movie, error) {
	row := q.db.QueryRowContext(ctx, getMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.ReleaseDate,
		&i.PosterUrl,
	)
	return i, err
}

const getMovies = `-- name: GetMovies :many
SELECT id, title, overview, release_date, poster_url FROM movies
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetMoviesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetMovies(ctx context.Context, arg GetMoviesParams) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, getMovies, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
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
