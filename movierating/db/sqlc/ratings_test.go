package db

import (
	"context"
	"testing"

	"github.com/aguazul-marco/hackbright/movierating/util"
)

func createRandomRating(t *testing.T) {
	arg := CreateRatingParams{
		Score: util.RandomRating(),
	}

	rating, err := testDB.CreateRating(context.Background(), arg)
	if err != nil {
		t.Errorf("error creating rating: %v", err)
	}

	if rating.Score == 0 {
		t.Errorf("got %v, wanted %v", rating.Score, arg.Score)
	}
}
