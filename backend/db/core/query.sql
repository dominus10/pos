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

-- name: RegisterRestaurantOwner :one
INSERT INTO employee (restaurant_id, role_id, name, email, password_hash) 
VALUES (
  NULL,
  (SELECT id FROM role WHERE role.name = 'admin'), 
  $1::TEXT,
  $2::TEXT,
  crypt($3, gen_salt('bf'))
)
RETURNING *;

-- name: AddEmployee :one
INSERT INTO employee (restaurant_id, role_id, name, email, password_hash) 
VALUES (
  (SELECT id FROM restaurant WHERE restaurant.name = $1),
  (SELECT id FROM role WHERE role.name = $2), 
  $3,
  $4,
  crypt($5, gen_salt('bf'))
)
RETURNING *;

-- name: GetEmployeeByEmail :one
SELECT * FROM employee
WHERE email = $1;

-- name: EmployeeClockIn :one
UPDATE employee
SET clock_in_time = NOW()
WHERE email = $1 AND password_hash = crypt($2, gen_salt('bf'))
AND clock_in_time IS NULL -- Prevent duplicate clock-ins
RETURNING *;

-- name: EmployeeClockOut :one
UPDATE employee
SET clock_out_time = NOW()
WHERE email = $1
AND clock_in_time IS NOT NULL -- Ensure employee has clocked in first
AND clock_out_time IS NULL -- Prevent multiple clock-outs
RETURNING *;

-- name: DelegateRoles :one
UPDATE employee
SET role_id = (SELECT id FROM role WHERE role.name = 'cashier')
WHERE email = $2
AND restaurant_id = (SELECT id FROM restaurant WHERE restaurant.name = $1)
RETURNING *;

-- name: GetAllEmployee :many
SELECT e.id, e.name, e.email, r.name AS role
FROM employee e
JOIN role r ON e.role_id = r.id
WHERE e.restaurant_id = (SELECT id FROM restaurant WHERE restaurant.name = $1);

-- name: GetEmployeeWorkHours :many
SELECT id, name, clock_in_time, clock_out_time,
       EXTRACT(EPOCH FROM (clock_out_time - clock_in_time)) / 3600 AS hours_worked
FROM employee
WHERE restaurant_id = $1::UUID
AND clock_in_time IS NOT NULL
AND clock_out_time IS NOT NULL
ORDER BY clock_in_time DESC;

-- name: RemoveEmployee :one
DELETE FROM employee 
WHERE email = $2
AND restaurant_id = (SELECT id FROM restaurant WHERE restaurant.name = $1)
RETURNING *;

-- name: TransferOwnership :one
UPDATE employee 
SET role_id = (SELECT id FROM role WHERE role.name = 'admin')
WHERE email = $2
AND restaurant_id = (SELECT id FROM restaurant WHERE restaurant.name = $1)
RETURNING *;

