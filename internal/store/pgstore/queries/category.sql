-- name: CreateCategory :execresult
INSERT INTO categories ("name")
VALUES ($1)
RETURNING id;

-- name: GetCategoryByName :one
SELECT id, name
FROM categories
WHERE name = $1;

-- name: DeleteCategory :execresult
DELETE categories
WHERE category_id = $1;