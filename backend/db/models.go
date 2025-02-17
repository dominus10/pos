// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Employee struct {
	ID           pgtype.UUID
	RestaurantID pgtype.UUID
	RoleID       pgtype.UUID
	Name         string
	Email        string
	PasswordHash string
	ClockInTime  pgtype.Timestamp
	ClockOutTime pgtype.Timestamp
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type Inventory struct {
	ID            pgtype.UUID
	RestaurantID  pgtype.UUID
	MenuItemID    pgtype.UUID
	StockQuantity pgtype.Int4
	RestockDate   pgtype.Timestamp
}

type MenuCategory struct {
	ID           pgtype.UUID
	RestaurantID pgtype.UUID
	Name         string
	Description  pgtype.Text
}

type MenuItem struct {
	ID            pgtype.UUID
	RestaurantID  pgtype.UUID
	CategoryID    pgtype.UUID
	Name          string
	Description   pgtype.Text
	Price         pgtype.Numeric
	StockQuantity pgtype.Int4
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
}

type OrderItem struct {
	ID         pgtype.UUID
	OrderID    pgtype.UUID
	MenuItemID pgtype.UUID
	Quantity   int32
	Price      pgtype.Numeric
}

type OrderList struct {
	ID           pgtype.UUID
	RestaurantID pgtype.UUID
	TableID      pgtype.UUID
	EmployeeID   pgtype.UUID
	Status       string
	TotalPrice   pgtype.Numeric
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type Restaurant struct {
	ID        pgtype.UUID
	Name      string
	Address   string
	CreatedAt pgtype.Timestamp
}

type Role struct {
	ID          pgtype.UUID
	Name        string
	Permissions []byte
}

type Table struct {
	ID           pgtype.UUID
	RestaurantID pgtype.UUID
	TableNumber  int32
	Status       string
}

type Transaction struct {
	ID              pgtype.UUID
	RestaurantID    pgtype.UUID
	OrderID         pgtype.UUID
	Amount          pgtype.Numeric
	PaymentMethod   string
	TransactionTime pgtype.Timestamp
}
