package database

import (
	"bootcamp_v1/internal/database/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)


func createRandomCourse(t *testing.T) (Courses, error){
  
  user, err := CreateRandomUser(t)
  require.NoError(t, err)

  bootcamp, err := CreateRandomBootcamp(t)
  require.NoError(t, err)


  arg := CreateCourseParams{
    Title : "Front End Web Development",
    Description: "This course will provide you with all of the essentials to become a successful frontend web developer. You will learn to master HTML, CSS and front end JavaScript, along with tools like Git, VSCode and front end frameworks like Vue",
    Weeks: "8",
    Tuition: "8000",
    MinimumSkill: MinimumSkill(util.RandomMinimumSkill()),
    ScholarshipAvailable:true,
    BootcampID: bootcamp.ID,
    UserID: user.ID,
  }

  course, err := TestQueries.CreateCourse(context.Background(), arg)
  require.NoError(t, err)
  require.NotEmpty(t, course)


  require.Equal(t, arg.Title, course.Title)
  require.Equal(t, arg.Description, course.Description)
  require.Equal(t, arg.Weeks, course.Weeks)
  require.Equal(t, arg.Tuition, course.Tuition)
  require.Equal(t, arg.MinimumSkill, course.MinimumSkill)
  require.Equal(t, arg.ScholarshipAvailable, course.ScholarshipAvailable)
  require.Equal(t, arg.BootcampID, course.BootcampID)
  require.Equal(t, arg.UserID, course.UserID)

  require.NotZero(t, course.ID)
  require.NotZero(t, course.CreatedAt)

  return course, nil
}

func TestCreateCourse(t *testing.T){
  createRandomCourse(t)
}


func TestGetCourse(t *testing.T){
  existingCourse, err := createRandomCourse(t)
  require.NoError(t, err)
  
  course, err := TestQueries.GetCourse(context.Background(), existingCourse.ID)
  require.NoError(t, err)
  require.NotEmpty(t, course)

  require.Equal(t, existingCourse.ID, course.ID)
  require.Equal(t, existingCourse.Title, course.Title)
  require.Equal(t, existingCourse.Description, course.Description)
  require.Equal(t, existingCourse.Weeks, course.Weeks)
  require.Equal(t, existingCourse.Tuition, course.Tuition)
  require.Equal(t, existingCourse.MinimumSkill, course.MinimumSkill)
  require.Equal(t, existingCourse.ScholarshipAvailable, course.ScholarshipAvailable)
  require.Equal(t, existingCourse.BootcampID, course.BootcampID)
  require.Equal(t, existingCourse.UserID, course.UserID)
  
  require.WithinDuration(t, existingCourse.CreatedAt.Time, course.CreatedAt.Time, time.Millisecond)

}


func TestUpdateCourse(t *testing.T){
  originalCourse, err := createRandomCourse(t)
  require.NoError(t, err)
  args := UpdateCourseParams{
    ID : originalCourse.ID,
    Title : originalCourse.Title,
    Description: originalCourse.Description,
    Weeks: originalCourse.Weeks,
    Tuition: originalCourse.Tuition,
    MinimumSkill: originalCourse.MinimumSkill,
    ScholarshipAvailable: originalCourse.ScholarshipAvailable,
  }
  updatedCourse, err := TestQueries.UpdateCourse(context.Background(), args)
  require.NoError(t, err)
  require.NotEmpty(t, updatedCourse)

  require.Equal(t, args.ID, updatedCourse.ID)
  require.Equal(t, args.Title, updatedCourse.Title)
  require.Equal(t, args.Description, updatedCourse.Description)
  require.Equal(t, args.Weeks, updatedCourse.Weeks)
  require.Equal(t, args.Tuition, updatedCourse.Tuition)
  require.Equal(t, args.MinimumSkill, updatedCourse.MinimumSkill)
  require.Equal(t, args.ScholarshipAvailable, updatedCourse.ScholarshipAvailable)

  require.True(t, updatedCourse.CreatedAt.Valid)
}


func TestDeleteCourse(t *testing.T){
  course, err := createRandomCourse(t)
  require.NoError(t,err)

  err = TestQueries.DeleteCourse(context.Background(), course.ID)
  require.NoError(t, err)
}
