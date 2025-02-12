-- name: CreateBootcamp :one
INSERT INTO bootcamps (
  user_id, name, slug, description, website, phone, email, address,careers, job_assistance, job_guarantee, accept_gi
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;


-- name: ListBootcamps :many
SELECT * FROM bootcamps
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetBootcamp :one
SELECT * FROM bootcamps
WHERE id = $1 LIMIT 1;

-- name: UpdateBootcamp :one
UPDATE bootcamps
  SET name = $2,
      slug = $3,
      description = $4,
      website = $5,
      phone = $6,
      email = $7,
      address = $8,
      careers = $9,
      job_assistance = $10,
      job_guarantee = $11,
      accept_gi = $12
WHERE id = $1
RETURNING *;

-- name: DeleteBootcamp :exec
DELETE FROM bootcamps
WHERE id = $1;
