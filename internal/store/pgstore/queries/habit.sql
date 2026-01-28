-- name: CreateHabit :one
INSERT INTO habits (
    "name",
    "category",
    "description",
    "frequency",
    "start_date",
    "target_date",
    "priority"
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: GetHabitById :one
SELECT * FROM habits WHERE id = $1;

-- name: GetHabitByName :one
SELECT * FROM habits WHERE name = $1;

-- name: UpdateHabit :one
UPDATE habits
SET
  name        = COALESCE(sqlc.narg(name), name),
  category    = COALESCE(sqlc.narg(category), category),
  description = COALESCE(sqlc.narg(description), description),
  frequency   = COALESCE(sqlc.narg(frequency), frequency),
  start_date  = COALESCE(sqlc.narg(start_date), start_date),
  target_date = COALESCE(sqlc.narg(target_date), target_date),
  priority    = COALESCE(sqlc.narg(priority), priority),
  updated_at  = NOW()
WHERE id = sqlc.arg(id)
RETURNING
id,
name,
category,
description,
frequency,
start_date,
target_date,
priority,
created_at,
updated_at;

-- name: DeleteHabit :execresult
DELETE FROM habits
WHERE id = $1;