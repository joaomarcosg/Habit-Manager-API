-- name: CreateCategory :one
INSERT INTO categories ("name")
VALUES ($1)
RETURNING id;

-- name: GetCategoryByName :one
SELECT id, name, entries, created_at, updated_at
FROM categories
WHERE name = $1;

-- name: GetCategoryEntries :one
SELECT entries
FROM categories
WHERE name = $1;

-- name: DeleteCategory :execresult
DELETE FROM categories
WHERE name = $1
AND entries = 0;