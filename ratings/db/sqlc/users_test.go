package db

import (
	"context"
	"testing"

	"github.com/phillipwright7/hackbright/ratings/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomString(10),
		Password: util.RandomString(20),
		Email:    util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	if err != nil {
		t.Error(err)
	}
	if user.Username != arg.Username {
		t.Errorf("username is not equal; want: %v, got: %v", arg.Username, user.Username)
	}
	if user.Password != arg.Password {
		t.Errorf("password is not equal; want: %v, got: %v", arg.Password, user.Password)
	}
	if user.Email != arg.Email {
		t.Errorf("email is not equal; want: %v, got: %v", arg.Email, user.Email)
	}

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestDeleteUser(t *testing.T) {
	arg := createRandomUser(t)

	if err := testQueries.DeleteUser(context.Background(), arg.Username); err != nil {
		t.Errorf("delete user failed: %v", arg.Username)
	}

}
