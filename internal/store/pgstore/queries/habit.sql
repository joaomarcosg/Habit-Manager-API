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
    name = COALESCE($2, name),
    category = COALESCE($3, category),
    description = COALESCE($4, description),
    frequency = COALESCE($5, frequency),
    start_date = COALESCE($6, start_date),
    target_date = COALESCE($7, target_date),
    priority = COALESCE($8, priority),
    updated_at = NOW()
WHERE id = $1
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