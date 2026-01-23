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

-- name: UpdateHabit :execresult
UPDATE habits SET (
    "name",
    "category",
    "description",
    "frequency",
    "start_date",
    "target_date",
    "priority"
)
WHERE id = $1;

-- name: DeleteHabit :execresult
DELETE FROM habits
WHERE id = $1;