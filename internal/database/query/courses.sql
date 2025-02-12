-- name: CreateCourse :one
INSERT INTO courses (
  title, description, weeks, tuition, minimum_skill, scholarship_available,bootcamp_id, user_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)RETURNING *;

-- name: ListCourses :many
SELECT * FROM courses
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetCourse :one
SELECT * FROM courses
WHERE id = $1 LIMIT 1;

-- name: UpdateCourse :one
UPDATE courses
  SET title = $2,
      description = $3,
      weeks = $4,
      tuition = $5,
      minimum_skill = $6,
      scholarship_available = $7
WHERE id = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses
WHERE id = $1;
