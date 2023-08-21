package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wizlif/tempmee_assignment/util"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByEmail(t *testing.T) {
	user := createRandomUser(t)

	userByEmail, err := testStore.GetUserByEmail(context.Background(), user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, userByEmail)

	require.Equal(t, user.Email, userByEmail.Email)
	require.Equal(t, user.Password, userByEmail.Password)
	require.Equal(t, user.CreatedAt, userByEmail.CreatedAt)
}

func createRandomUser(t *testing.T) User {

	arg := CreateUserParams{
		Email:    util.RandomEmail(),
		Password: util.RandomString(8),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.CreatedAt)
	require.NoError(t, err)

	return user
}
