// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package tenant

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTenant = `-- name: CreateTenant :one
INSERT INTO tenants (name, email, tier, db_name, db_host, db_usage)
VALUES ($1, $2, 'free', $3, $4, 0)
RETURNING id, name, email, tier, db_usage, db_name, db_host, created_at
`

type CreateTenantParams struct {
	Name   string
	Email  string
	DbName pgtype.Text
	DbHost pgtype.Text
}

func (q *Queries) CreateTenant(ctx context.Context, arg CreateTenantParams) (Tenant, error) {
	row := q.db.QueryRow(ctx, createTenant,
		arg.Name,
		arg.Email,
		arg.DbName,
		arg.DbHost,
	)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Tier,
		&i.DbUsage,
		&i.DbName,
		&i.DbHost,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTenant = `-- name: DeleteTenant :exec
DELETE FROM tenants WHERE id = $1
`

func (q *Queries) DeleteTenant(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteTenant, id)
	return err
}

const downgradeToFreeUser = `-- name: DowngradeToFreeUser :one
UPDATE tenants SET tier = 'free' WHERE id = $1
RETURNING id, name, email, tier, db_usage, db_name, db_host, created_at
`

func (q *Queries) DowngradeToFreeUser(ctx context.Context, id pgtype.UUID) (Tenant, error) {
	row := q.db.QueryRow(ctx, downgradeToFreeUser, id)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Tier,
		&i.DbUsage,
		&i.DbName,
		&i.DbHost,
		&i.CreatedAt,
	)
	return i, err
}

const getPaidTenants = `-- name: GetPaidTenants :many
SELECT id, name, email, tier, db_usage, db_name, db_host, created_at FROM tenants WHERE tier = 'paid'
`

func (q *Queries) GetPaidTenants(ctx context.Context) ([]Tenant, error) {
	rows, err := q.db.Query(ctx, getPaidTenants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tenant
	for rows.Next() {
		var i Tenant
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Tier,
			&i.DbUsage,
			&i.DbName,
			&i.DbHost,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTenantByID = `-- name: GetTenantByID :one
SELECT id, name, email, tier, db_usage, db_name, db_host, created_at FROM tenants WHERE id = $1
`

func (q *Queries) GetTenantByID(ctx context.Context, id pgtype.UUID) (Tenant, error) {
	row := q.db.QueryRow(ctx, getTenantByID, id)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Tier,
		&i.DbUsage,
		&i.DbName,
		&i.DbHost,
		&i.CreatedAt,
	)
	return i, err
}

const updateTenantDBUsage = `-- name: UpdateTenantDBUsage :exec
UPDATE tenants SET db_usage = db_usage + $1 WHERE id = $2
`

type UpdateTenantDBUsageParams struct {
	DbUsage pgtype.Int8
	ID      pgtype.UUID
}

func (q *Queries) UpdateTenantDBUsage(ctx context.Context, arg UpdateTenantDBUsageParams) error {
	_, err := q.db.Exec(ctx, updateTenantDBUsage, arg.DbUsage, arg.ID)
	return err
}

const upgradeToPaidUser = `-- name: UpgradeToPaidUser :one
UPDATE tenants SET tier = 'paid' WHERE id = $1 
RETURNING id, name, email, tier, db_usage, db_name, db_host, created_at
`

func (q *Queries) UpgradeToPaidUser(ctx context.Context, id pgtype.UUID) (Tenant, error) {
	row := q.db.QueryRow(ctx, upgradeToPaidUser, id)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Tier,
		&i.DbUsage,
		&i.DbName,
		&i.DbHost,
		&i.CreatedAt,
	)
	return i, err
}
