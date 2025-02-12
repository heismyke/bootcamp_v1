package database

import (
	"bootcamp_v1/internal/database/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)


func CreateRandomBootcamp(t *testing.T) (Bootcamps, error) {
  user1, err := CreateRandomUser(t)
  require.NoError(t, err)
  
  params := CreateBootcampParams{
    UserID : sql.NullInt64{
      Int64 : user1.ID,
      Valid : true,
    },
    Name : "Devworks Bootcamp", 
    Slug : sql.NullString{
      String : util.RandomSlug(),
      Valid : true,
    },
    Description: "Devworks is a full stack JavaScript Bootcamp located in the heart of Boston that focuses on the technologies you need to get a high paying job as a web developer",
    Website: "https://devworks.com",
    Phone: "(111)111-1111",
    Email: "enroll@devworks.com",
    Address: "233 Bay State Rd Boston MA 02215",
    Careers: []byte(`["Web Development", "UI/UX", "Business"]`),
    JobAssistance: true,
    JobGuarantee: false,
    AcceptGi:true,
    
  }

  bootcamp, err := TestQueries.CreateBootcamp(context.Background(), params)
  require.NoError(t, err)
  require.NotEmpty(t, bootcamp)

  require.Equal(t, params.UserID, bootcamp.UserID)
  require.Equal(t, params.Slug, bootcamp.Slug)
  require.Equal(t, params.Description, bootcamp.Description)
  require.Equal(t, params.Website, bootcamp.Website)
  require.Equal(t, params.Phone, bootcamp.Phone)
  require.Equal(t, params.Email, bootcamp.Email)
  require.Equal(t, params.Address, bootcamp.Address)
  require.Equal(t, params.Careers, bootcamp.Careers)
  require.Equal(t, params.JobAssistance, bootcamp.JobAssistance)
  require.Equal(t, params.JobGuarantee, bootcamp.JobGuarantee)
  require.Equal(t, params.AcceptGi, bootcamp.AcceptGi)
  
  require.NotZero(t, bootcamp.ID)
  require.NotZero(t, bootcamp.CreatedAt)


  return bootcamp, nil
}



func TestCreateBootcamp(t *testing.T){
  CreateRandomBootcamp(t)
}



func TestGetBootcamp(t *testing.T){
  
  existingBootcamp, err := CreateRandomBootcamp(t)
  require.NoError(t, err)
  
  bootcamp, err := TestQueries.GetBootcamp(context.Background(), existingBootcamp.ID)
  require.NoError(t, err)
  require.NotEmpty(t, bootcamp)
  
  require.Equal(t, existingBootcamp.ID, bootcamp.ID)
  require.Equal(t, existingBootcamp.UserID, bootcamp.UserID)
  require.Equal(t, existingBootcamp.Name, bootcamp.Name)
  require.Equal(t, existingBootcamp.Slug, bootcamp.Slug)
  require.Equal(t, existingBootcamp.Description, bootcamp.Description)
  require.Equal(t, existingBootcamp.Website, bootcamp.Website)
  require.Equal(t, existingBootcamp.Phone, bootcamp.Phone)
  require.Equal(t, existingBootcamp.Email, bootcamp.Email)
  require.Equal(t, existingBootcamp.Address, bootcamp.Address)
  require.Equal(t, existingBootcamp.Careers, bootcamp.Careers)
  require.Equal(t, existingBootcamp.JobAssistance, bootcamp.JobAssistance)
  require.Equal(t, existingBootcamp.JobGuarantee, bootcamp.JobGuarantee)
  require.Equal(t, existingBootcamp.AcceptGi, bootcamp.AcceptGi)
  require.WithinDuration(t, existingBootcamp.CreatedAt.Time, bootcamp.CreatedAt.Time, time.Millisecond)
  
}


func TestUpdateBootcamp(t *testing.T){
  
  originalBootcamp, err := CreateRandomBootcamp(t)
  require.NoError(t, err)

  args := UpdateBootcampParams{
    ID : originalBootcamp.ID,
    Name : originalBootcamp.Name,
    Slug: originalBootcamp.Slug,
    Description: originalBootcamp.Description,
    Website: originalBootcamp.Website,
    Phone: originalBootcamp.Phone,
    Email: originalBootcamp.Email,
    Address: originalBootcamp.Address,
    Careers: originalBootcamp.Careers,
    JobAssistance: originalBootcamp.JobAssistance,
    JobGuarantee: originalBootcamp.JobGuarantee,
    AcceptGi:originalBootcamp.AcceptGi,
  }

  updatedBootcamp, err := TestQueries.UpdateBootcamp(context.Background(), args)
  require.NoError(t, err)
  require.NotEmpty(t, updatedBootcamp)
  
  require.Equal(t, args.ID, updatedBootcamp.ID)
  require.Equal(t, args.Name, updatedBootcamp.Name)
  require.Equal(t, args.Slug, updatedBootcamp.Slug)
  require.Equal(t, args.Description, updatedBootcamp.Description)
  require.Equal(t, args.Website, updatedBootcamp.Website)
  require.Equal(t, args.Phone, updatedBootcamp.Phone)
  require.Equal(t, args.Email, updatedBootcamp.Email)
  require.Equal(t, args.Address, updatedBootcamp.Address)
  require.Equal(t, args.Careers, updatedBootcamp.Careers)
  require.Equal(t, args.JobAssistance, updatedBootcamp.JobAssistance)
  require.Equal(t, args.JobGuarantee, updatedBootcamp.JobGuarantee)
  require.Equal(t, args.AcceptGi, updatedBootcamp.AcceptGi)
  
  require.True(t, updatedBootcamp.CreatedAt.Valid)
}


func TestDeleteBootcamp(t *testing.T){
  bootcamp, err := CreateRandomBootcamp(t)
  require.NoError(t, err)

   err = TestQueries.DeleteBootcamp(context.Background(), bootcamp.ID)
  require.NoError(t, err)

}
