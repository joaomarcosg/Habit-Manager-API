-- name: CreateCategory :execresult
INSERT INTO categories ("name")
VALUES ($1)
RETURNING id;

-- name: GetCategoryById :one
SELECT id, name, entries, created_at, updated_at
FROM categories
WHERE id = $1;

-- name: GetCategoryByName :one
SELECT id, name, entries, created_at, updated_at
FROM categories
WHERE name = $1;

-- name: DeleteCategory :execresult
DELETE FROM categories
WHERE category_id = $1;