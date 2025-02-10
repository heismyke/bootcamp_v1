-- name: CreateBootcamp :one
INSERT INTO bootcamps (
  user_id, name, slug, description, website, phone, email, address, latitude, longitude, location_details, careers, average_rating, average_cost, photo, housing, job_assistance, job_guarantee, accept_gi
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
) RETURNING id, created_at;


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
  SET user_id = $2,
      name = $3,
      slug = $4,
      description = $5,
      website = $6,
      phone = $7,
      email = $8,
      address = $9,
      latitude = $10,
      longitude = $11,
      location_details = $12,
      careers = $13,
      average_rating = $14,
      average_cost = $15,
      photo = $16,
      housing = $17,
      job_assistance = $18,
      job_guarantee = $19,
      accept_gi = $20
WHERE id = $1
RETURNING *;

-- name: DeleteBootcamp :exec
DELETE FROM bootcamps
WHERE id = $1;
