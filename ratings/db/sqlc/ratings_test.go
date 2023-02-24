package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/phillipwright7/hackbright/ratings/util"
)

func createRandomRating(t *testing.T) Rating {
	arg := CreateRatingParams{
		MovieID: sql.NullInt32{},
		UserID:  sql.NullInt32{},
		Score:   util.RandomInt(1, 5),
	}

	rating, err := testQueries.CreateRating(context.Background(), arg)

	if err != nil {
		t.Error(err)
	}
	if rating.MovieID != arg.MovieID {
		t.Errorf("movie id is not equal; want: %v, got: %v", arg.MovieID, rating.MovieID)
	}
	if rating.UserID != arg.UserID {
		t.Errorf("user id is not equal; want: %v, got: %v", arg.UserID, rating.UserID)
	}
	if rating.Score != arg.Score {
		t.Errorf("score is not equal; want: %v, got: %v", arg.Score, rating.Score)
	}

	return rating
}

func TestCreateRating(t *testing.T) {
	createRandomRating(t)
}

func TestDeleteRating(t *testing.T) {
	arg := createRandomRating(t)
	argParams := GetRatingDetailsParams{
		MovieID: arg.MovieID,
		UserID:  arg.UserID,
	}

	if err := testQueries.DeleteRating(context.Background(), DeleteRatingParams(argParams)); err != nil {
		t.Errorf("delete rating failed: %v", arg.MovieID)
	}
}

func TestGetMovieRatings(t *testing.T) {
	arg := createRandomRating(t)

	ratings, err := testQueries.GetMovieRatings(context.Background(), arg.MovieID)

	if err != nil {
		t.Error(err)
	}
	for _, r := range ratings {
		if r.RatingID == arg.RatingID {
			break
		}
		t.Errorf("can't get ratings; want %v, got: %v", arg.RatingID, r.RatingID)
	}
}

func TestGetRatingDetails(t *testing.T) {
	arg := createRandomRating(t)
	argParams := GetRatingDetailsParams{
		MovieID: arg.MovieID,
		UserID:  arg.UserID,
	}

	rating, err := testQueries.GetRatingDetails(context.Background(), argParams)

	log.Println(arg)
	log.Println(rating)

	if err != nil {
		t.Error(err)
	}
	if rating.MovieID != arg.MovieID {
		t.Errorf("movie id is not equal; want: %v, got: %v", arg.MovieID, rating.MovieID)
	}
	if rating.UserID != arg.UserID {
		t.Errorf("user id is not equal; want: %v, got: %v", arg.UserID, rating.UserID)
	}
	if rating.Score != arg.Score {
		t.Errorf("score is not equal; want: %v, got: %v", arg.Score, rating.Score)
	}
	if rating.RatingID != arg.RatingID {
		t.Errorf("rating id is not equal; want: %v, got: %v", arg.UserID, rating.UserID)
	}
}
