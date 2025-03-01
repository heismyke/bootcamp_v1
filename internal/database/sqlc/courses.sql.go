// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: courses.sql

package database

import (
	"context"
)

const createCourse = `-- name: CreateCourse :one
INSERT INTO courses (
  title, description, weeks, tuition, minimum_skill, scholarship_available,bootcamp_id, user_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)RETURNING id, title, description, weeks, tuition, minimum_skill, scholarship_available, bootcamp_id, user_id, created_at
`

type CreateCourseParams struct {
	Title                string       `json:"title"`
	Description          string       `json:"description"`
	Weeks                string       `json:"weeks"`
	Tuition              string       `json:"tuition"`
	MinimumSkill         MinimumSkill `json:"minimum_skill"`
	ScholarshipAvailable bool         `json:"scholarship_available"`
	BootcampID           int64        `json:"bootcamp_id"`
	UserID               int64        `json:"user_id"`
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Courses, error) {
	row := q.db.QueryRowContext(ctx, createCourse,
		arg.Title,
		arg.Description,
		arg.Weeks,
		arg.Tuition,
		arg.MinimumSkill,
		arg.ScholarshipAvailable,
		arg.BootcampID,
		arg.UserID,
	)
	var i Courses
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Weeks,
		&i.Tuition,
		&i.MinimumSkill,
		&i.ScholarshipAvailable,
		&i.BootcampID,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCourse = `-- name: DeleteCourse :exec
DELETE FROM courses
WHERE id = $1
`

func (q *Queries) DeleteCourse(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCourse, id)
	return err
}

const getCourse = `-- name: GetCourse :one
SELECT id, title, description, weeks, tuition, minimum_skill, scholarship_available, bootcamp_id, user_id, created_at FROM courses
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCourse(ctx context.Context, id int64) (Courses, error) {
	row := q.db.QueryRowContext(ctx, getCourse, id)
	var i Courses
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Weeks,
		&i.Tuition,
		&i.MinimumSkill,
		&i.ScholarshipAvailable,
		&i.BootcampID,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const listCourses = `-- name: ListCourses :many
SELECT id, title, description, weeks, tuition, minimum_skill, scholarship_available, bootcamp_id, user_id, created_at FROM courses
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCoursesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCourses(ctx context.Context, arg ListCoursesParams) ([]Courses, error) {
	rows, err := q.db.QueryContext(ctx, listCourses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Courses
	for rows.Next() {
		var i Courses
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Weeks,
			&i.Tuition,
			&i.MinimumSkill,
			&i.ScholarshipAvailable,
			&i.BootcampID,
			&i.UserID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCourse = `-- name: UpdateCourse :one
UPDATE courses
  SET title = $2,
      description = $3,
      weeks = $4,
      tuition = $5,
      minimum_skill = $6,
      scholarship_available = $7
WHERE id = $1
RETURNING id, title, description, weeks, tuition, minimum_skill, scholarship_available, bootcamp_id, user_id, created_at
`

type UpdateCourseParams struct {
	ID                   int64        `json:"id"`
	Title                string       `json:"title"`
	Description          string       `json:"description"`
	Weeks                string       `json:"weeks"`
	Tuition              string       `json:"tuition"`
	MinimumSkill         MinimumSkill `json:"minimum_skill"`
	ScholarshipAvailable bool         `json:"scholarship_available"`
}

func (q *Queries) UpdateCourse(ctx context.Context, arg UpdateCourseParams) (Courses, error) {
	row := q.db.QueryRowContext(ctx, updateCourse,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Weeks,
		arg.Tuition,
		arg.MinimumSkill,
		arg.ScholarshipAvailable,
	)
	var i Courses
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Weeks,
		&i.Tuition,
		&i.MinimumSkill,
		&i.ScholarshipAvailable,
		&i.BootcampID,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}
