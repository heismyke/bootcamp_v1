package database

import (
	"bootcamp_v1/internal/database/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)


func CreateRandomUser(t *testing.T) (Users, error) {
  plainTextPassword := util.RandomPassword()
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)
  require.NoError(t, err)
	params := CreateUserParams{
		Name:     util.RandomName(),
		Email:    util.RandomEmail(),
		Role:     UserRole(util.RandomRole()),
		Password: string(hashedPassword),
	}

	createdUser, err := TestQueries.CreateUser(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, createdUser)

	require.Equal(t, params.Name, createdUser.Name)
	require.Equal(t, params.Email, createdUser.Email)
	require.Equal(t, params.Role, createdUser.Role)
  require.Equal(t, params.Password, createdUser.Password)

	require.NotZero(t, createdUser.ID)
	require.NotZero(t, createdUser.CreatedAt)

	user := Users{
		ID:        createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		Role:      createdUser.Role,
		Password:  plainTextPassword,
		CreatedAt: createdUser.CreatedAt,
	}

	return user, nil
}


func TestCreateUser(t *testing.T){
  CreateRandomUser(t)  
}

func TestGetUser(t *testing.T){
  
  user1, err := CreateRandomUser(t)
  require.NoError(t, err)

  user2, err := TestQueries.GetUser(context.Background(), user1.ID)
  require.NoError(t, err)
  require.NotEmpty(t, user2)

  require.Equal(t, user1.ID, user2.ID)
  require.Equal(t, user1.Name, user2.Name)
  require.Equal(t, user1.Email, user2.Email)
  require.Equal(t, user1.Role, user2.Role)
  err = bcrypt.CompareHashAndPassword([]byte(user2.Password), []byte(user1.Password))
  require.NoError(t, err)
  require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Millisecond) 
 
}



func TestUpdateUser(t *testing.T) {
    originalUser, err := CreateRandomUser(t)
    require.NoError(t, err)

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originalUser.Password), bcrypt.DefaultCost)
    require.NoError(t, err)

    args := UpdateUserParams{
        ID:       originalUser.ID,
        Name:     originalUser.Name,
        Email:    originalUser.Email,
        Role:     originalUser.Role,
        Password: string(hashedPassword),  
    }

    updatedUser, err := TestQueries.UpdateUser(context.Background(), args)
    require.NoError(t, err)
    require.NotEmpty(t, updatedUser)

    require.Equal(t, originalUser.ID, updatedUser.ID)
    require.Equal(t, originalUser.Name, updatedUser.Name)
    require.Equal(t, originalUser.Email, updatedUser.Email)
    require.Equal(t, originalUser.Role, updatedUser.Role)

    // Compare the hashed password (updatedUser.Password) with the plain password (originalUser.Password)
    err = bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte(originalUser.Password))
    require.NoError(t, err)  // This ensures the passwords match

    require.True(t, updatedUser.CreatedAt.Valid)
}



func TestDeleteUser(t *testing.T) {
    user1, err := CreateRandomUser(t)
    require.NoError(t, err)

    err = TestQueries.DeleteUser(context.Background(), user1.ID)
    require.NoError(t, err)

    user2, err := TestQueries.GetUser(context.Background(), user1.ID)
    require.Error(t, err)  
    require.Empty(t, user2) 
}
