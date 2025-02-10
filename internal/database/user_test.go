package database

import (
	"context"
	"testing"
  sqlc "bootcamp_v1/internal/database/sqlc"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount (t *testing.T){
  arg := sqlc.CreateUserParams{
    Name : "User Account",
    Email : "useraccount@gmail.com",
    Role : "user",
    Password : "12345",
  } 

  user, err := TestQueries.CreateUser(context.Background(), arg)
    require.NoError(t,err)
    require.NotEmpty(t, user)

  require.Equal(t, arg.Name, user.Name)
  require.Equal(t, arg.Email, user.Email)
  require.Equal(t, arg.Role, user.Role)
  require.Equal(t, arg.Password, user.Password)

  require.NotZero(t, user.ID)
  require.NotZero(t, user.CreatedAt)
}
