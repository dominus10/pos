-- name: GetTenantByID :one
SELECT * FROM tenants WHERE id = $1;

-- name: CreateTenant :one
INSERT INTO tenants (name, email, tier, db_name, db_host, db_usage)
VALUES ($1, $2, 'free', $3, $4, 0)
RETURNING *;

-- name: UpgradeToPaidUser :one
UPDATE tenants SET tier = 'paid' WHERE id = $1 
RETURNING *;

-- name: GetPaidTenants :many
SELECT * FROM tenants WHERE tier = 'paid';

-- name: DowngradeToFreeUser :one
UPDATE tenants SET tier = 'free' WHERE id = $1
RETURNING *;

-- name: UpdateTenantDBUsage :exec
UPDATE tenants SET db_usage = db_usage + $1 WHERE id = $2;

-- name: DeleteTenant :exec
DELETE FROM tenants WHERE id = $1;