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

-- name:RegisterRestaurantOwner :one
INSERT INTO employee (restaurant_id, role_id, name, email, password_hash) 
VALUES (
    NULL, -- No restaurant assigned yet
    (SELECT id FROM role WHERE name = 'admin'), 
    $3,
    $4,
    crypt($5, gen_salt('bf'))
)
RETURNING *;

-- name:AddEmployee :one
INSERT INTO employee (restaurant_id, role_id, name, email, password_hash) 
VALUES (
    (SELECT id FROM restaurant WHERE name = $1),
    (SELECT id FROM role WHERE name = $2), 
    $3,
    $4,
    crypt($5, gen_salt('bf'))
)
RETURNING *;

-- name:DelegateRoles :one
UPDATE employee
SET role_id = (SELECT id FROM role WHERE name = 'cashier')
WHERE email = $2
AND restaurant_id = (SELECT id FROM restaurant WHERE name = $1)
RETURNING *;

-- name:GetAllEmployee :many
SELECT e.id, e.name, e.email, r.name AS role
FROM employee e
JOIN role r ON e.role_id = r.id
WHERE e.restaurant_id = (SELECT id FROM restaurant WHERE name = $1);

-- name:RemoveEmployee :one
DELETE FROM employee 
WHERE email = $2
AND restaurant_id = (SELECT id FROM restaurant WHERE name = $1)
RETURNING *;

-- name:TransferOwnership :one
UPDATE employee 
SET role_id = (SELECT id FROM role WHERE name = 'admin')
WHERE email = $2
AND restaurant_id = (SELECT id FROM restaurant WHERE name = $1)
RETURNING *;