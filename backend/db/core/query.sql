-- name: InsertNewRestaurant :one
INSERT INTO restaurant (name, address) 
VALUES ($1,$2)
RETURNING *;

-- name: UpdateExistingRestaurant :one
UPDATE restaurant 
SET name = $1, address = $2
WHERE id = $3
RETURNING *;

-- name: GetAllRestaurant :many
SELECT * FROM restaurant ORDER BY created_at DESC;

-- name: GetRestaurant :one
SELECT * FROM restaurant WHERE id = $1;

-- name: DeleteRestaurant :one
DELETE FROM restaurant 
WHERE id = $1
RETURNING *;