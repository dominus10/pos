package user

import (
	"context"

	"github.com/dominus10/pos/db"
)

func createFreeTenant(ctx context.Context, q *db.Queries, name, email string) (string, error) {
	var tenantID string
	// row := q.QueryRowContext(ctx, "INSERT INTO tenants (name, email, tier) VALUES ($1, $2, 'free') RETURNING id", name, email)
	// if err := row.Scan(&tenantID); err != nil {
	// 	return "", err
	// }
	return tenantID, nil
}