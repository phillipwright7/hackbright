package db

import (
	"context"
	"testing"

	"github.com/phillipwright7/hackbright/ratings/util"
)

func createRandomMovie(t *testing.T) Movie {
	arg := CreateMovieParams{
		Title:       util.RandomString(6),
		Overview:    util.RandomString(30),
		ReleaseDate: util.RandomDate(),
		PosterUrl:   util.RandomUrl(),
	}

	movie, err := testQueries.CreateMovie(context.Background(), arg)

	if err != nil {
		t.Error(err)
	}
	if movie.Title != arg.Title {
		t.Errorf("title is not equal; want: %v, got: %v", arg.Title, movie.Title)
	}
	if movie.Overview != arg.Overview {
		t.Errorf("overview is not equal; want: %v, got: %v", arg.Overview, movie.Overview)
	}
	if movie.ReleaseDate.Format("YYYY-MM-DD") != arg.ReleaseDate.Format("YYYY-MM-DD") {
		t.Errorf("release date is not equal; want: %v, got: %v", arg.ReleaseDate, movie.ReleaseDate)
	}
	if movie.PosterUrl != arg.PosterUrl {
		t.Errorf("poster url is not equal; want: %v, got: %v", arg.PosterUrl, movie.PosterUrl)
	}

	return movie
}

func TestCreateMovie(t *testing.T) {
	createRandomMovie(t)
}

func TestDeleteMovie(t *testing.T) {
	arg := createRandomMovie(t)

	if err := testQueries.DeleteMovie(context.Background(), arg.Title); err != nil {
		t.Errorf("delete movie failed: %v", arg.Title)
	}

}

func TestGetAllMovies(t *testing.T) {
	movies, err := testQueries.GetAllMovies(context.Background())

	if err != nil {
		t.Error(err)
	}
	if movies == nil {
		t.Errorf("movie list is nil: %v", err)
	}
}

func TestGetMovieDetails(t *testing.T) {
	arg := createRandomMovie(t)

	movie, err := testQueries.GetMovieDetails(context.Background(), arg.Title)

	if err != nil {
		t.Error(err)
	}
	if movie.Title != arg.Title {
		t.Errorf("title is not equal; want: %v, got: %v", arg.Title, movie.Title)
	}
	if movie.Overview != arg.Overview {
		t.Errorf("overview is not equal; want: %v, got: %v", arg.Overview, movie.Overview)
	}
	if movie.ReleaseDate != arg.ReleaseDate {
		t.Errorf("release date is zero; want: %v, got: %v", arg.ReleaseDate, movie.ReleaseDate)
	}
	if movie.PosterUrl != arg.PosterUrl {
		t.Errorf("poster url is not empty; want: %v, got: %v", arg.PosterUrl, movie.PosterUrl)
	}
}
