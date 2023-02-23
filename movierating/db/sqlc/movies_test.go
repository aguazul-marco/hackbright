package db

import (
	"context"
	"testing"

	"github.com/aguazul-marco/hackbright/movierating/util"
)

func createRandomMovie(t *testing.T) Movie {
	arg := CreateMovieParams{
		Title:       util.RandomString(6),
		Overview:    util.RandomString(32),
		ReleaseDate: util.RandomString(10),
		PosterUrl:   util.RandomString(10),
	}

	movie, err := testQueries.CreateMovie(context.Background(), arg)
	if err != nil {
		t.Errorf("error creating movie: %v", err)
	}

	if movie.Title == "" || movie.Overview == "" || movie.PosterUrl == "" || movie.ReleaseDate == "" {
		t.Errorf("empty movie")
	}

	switch {
	case movie.Title != arg.Title:
		t.Errorf("empty rating; got %v, wanted %v", movie.Title, arg.Title)
	case movie.Overview != arg.Overview:
		t.Errorf("empty rating; got %v, wanted %v", movie.Overview, arg.Overview)
	case movie.PosterUrl != arg.PosterUrl:
		t.Errorf("empty rating; got %v, wanted %v", movie.PosterUrl, arg.PosterUrl)
	case movie.ReleaseDate != arg.ReleaseDate:
		t.Errorf("empty rating; got %v, wanted %v", movie.ReleaseDate, arg.ReleaseDate)
	}

	return movie
}

func TestCreateMovie(t *testing.T) {
	createRandomMovie(t)
}

func TestGetMovie(t *testing.T) {
	movie1 := createRandomMovie(t)

	movie2, err := testQueries.GetMovie(context.Background(), movie1.ID)
	if err != nil {
		t.Errorf("error creating rating: %v", err)
	}

	var emptyMovie Movie
	if movie1 == emptyMovie {
		t.Errorf("no rating exist")
	}
	if movie1.ID != movie2.ID {
		t.Errorf("empty rating; got %v, wanted %v", movie1.ID, movie2.ID)
	}
}

func TestGetMovies(t *testing.T) {

	args := GetMoviesParams{
		Limit:  5,
		Offset: 0,
	}

	movieList, err := testQueries.GetMovies(context.Background(), args)
	if err != nil {
		t.Errorf("error listing users: got %v", err)
	}

	for _, movie := range movieList {
		if movie.Title == "" {
			t.Errorf("wanted movie title got %v", movie.Title)
		}
		if movie.Overview == "" {
			t.Errorf("wanted movie overview got %v", movie.Overview)
		}
		if movie.ReleaseDate == "" {
			t.Errorf("wanted release date got %v", movie.ReleaseDate)
		}
		if movie.PosterUrl == "" {
			t.Errorf("wanted poster url got %v", movie.PosterUrl)
		}

	}
}

func TestDeleteMovie(t *testing.T) {
	movie1 := createRandomMovie(t)
	err := testQueries.DeleteMovie(context.Background(), movie1.ID)
	if err != nil {
		t.Errorf("error occured while deleting: %v", err)
	}

	movie2, err := testQueries.GetRating(context.Background(), movie1.ID)
	if err == nil {
		t.Errorf("not empty: %v. got: %v", err, movie2.ID)
	}
}
