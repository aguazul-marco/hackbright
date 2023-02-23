package db

import (
	"context"
	"testing"

	"github.com/aguazul-marco/hackbright/movierating/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUser(),
		Password: util.RandomString(8),
		Email:    util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Errorf("error creating user: %v", err)
	}

	if user.Email == "" || user.Password == "" || user.Username == "" {
		t.Errorf("error: empty field")
	}

	switch {
	case user.Email != arg.Email:
		t.Errorf("error; got %v, wanted %v", user.Email, arg.Email)
	case user.Password != arg.Password:
		t.Errorf("empty rating; got %v, wanted %v", user.Password, arg.Password)
	case user.Username != arg.Username:
		t.Errorf("empty rating; got %v, wanted %v", user.Username, arg.Username)
	}

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	if err != nil {
		t.Errorf("error creating rating: %v", err)
	}

	var emptyUser User
	if user2 == emptyUser {
		t.Errorf("no rating exist")
	}
	if user1.ID != user2.ID {
		t.Errorf("empty rating; got %v, wanted %v", user1.ID, user2.ID)
	}
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	if err != nil {
		t.Errorf("error occured while deleting: %v", err)
	}

	user2, err := testQueries.GetRating(context.Background(), user1.ID)
	if err == nil {
		t.Errorf("not empty: %v. got: %v", err, user2.ID)
	}
}

func TestGetUsers(t *testing.T) {

	args := GetUsersParams{
		Limit:  5,
		Offset: 0,
	}

	userslist, err := testQueries.GetUsers(context.Background(), args)
	if err != nil {
		t.Errorf("error listing users: got %v", err)
	}

	if userslist == nil {
		t.Errorf("empty list")
	} else {
		for _, user := range userslist {
			if user.Username == "" {
				t.Errorf("wanted username got %v", user.Username)
			}
			if user.Password == "" {
				t.Errorf("wanted user password got %v", user.Password)
			}
			if user.Email == "" {
				t.Errorf("wanted email got %v", user.Email)
			}

		}
	}

}
