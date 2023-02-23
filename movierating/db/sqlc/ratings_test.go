package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/aguazul-marco/hackbright/movierating/util"
)

func createRandomRating(t *testing.T) Rating {
	user := createRandomUser(t)
	movie := createRandomMovie(t)

	userid := sql.NullInt32{Int32: int32(user.ID), Valid: true}
	movieid := sql.NullInt32{Int32: int32(movie.ID), Valid: true}

	arg := CreateRatingParams{
		Score:   util.RandomRating(),
		MovieID: movieid,
		UserID:  userid,
	}

	rating, err := testQueries.CreateRating(context.Background(), arg)
	if err != nil {
		t.Errorf("error creating rating: %v", err)
	}

	if rating.Score == 0 {
		t.Errorf("empty rating; got %v, wanted %v", rating.Score, arg.Score)
	}

	return rating
}

func TestCreateRating(t *testing.T) {
	createRandomRating(t)
}

func TestGetRating(t *testing.T) {
	rating1 := createRandomRating(t)

	rating2, err := testQueries.GetRating(context.Background(), rating1.ID)
	if err != nil {
		t.Errorf("error creating rating: %v", err)
	}

	var emptyRating Rating
	if rating2 == emptyRating {
		t.Errorf("no rating exist")
	}
	if rating1.ID != rating2.ID {
		t.Errorf("empty rating; got %v, wanted %v", rating1.ID, rating2.ID)
	}
}

func TestDeleteRating(t *testing.T) {
	rate1 := createRandomRating(t)
	err := testQueries.DeleteRating(context.Background(), rate1.ID)
	if err != nil {
		t.Errorf("error occured while deleting: %v", err)
	}

	rate2, err := testQueries.GetRating(context.Background(), rate1.ID)
	if err == nil {
		t.Errorf("not empty: %v. got: %v", err, rate2.ID)
	}
}

func TestGetMovieRatings(t *testing.T) {
	arg := createRandomRating(t)

	userRatings, err := testQueries.GetMovieRatings(context.Background(), arg.MovieID)
	if err != nil {
		t.Errorf("error retrieving movie rating: %v", err)
	}

	for _, rating := range userRatings {
		if rating.Score != arg.Score {
			t.Errorf("wanted: %v, got: %v", arg.Score, rating.Score)
		}
	}
}

func TestGetUserRatings(t *testing.T) {
	arg := createRandomRating(t)

	userRatings, err := testQueries.GetUserRatings(context.Background(), arg.UserID)
	if err != nil {
		t.Errorf("error retrieving user rating: %v", err)
	}

	for _, rating := range userRatings {
		if rating.Score != arg.Score {
			t.Errorf("wanted: %v, got: %v", arg.Score, rating.Score)
		}
	}
}

func TestUpdateRating(t *testing.T) {
	arg := createRandomRating(t)
	updateRating := UpdateRatingParams{
		ID:    arg.ID,
		Score: util.RandomRating(),
	}

	newRating, err := testQueries.UpdateRating(context.Background(), updateRating)
	if err != nil {
		t.Errorf("error updating rating: %v", err)
	}

	if newRating.Score == arg.Score {
		t.Errorf("want: %v, got: %v", newRating.Score, arg.Score)
	}

}
