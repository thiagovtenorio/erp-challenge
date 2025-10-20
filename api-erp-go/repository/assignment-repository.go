package repository

import (
	db "api-erp-go/db-connection"
	"api-erp-go/model"
	"context"
	"fmt"
	"log"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AssignDelivery(delivery model.Delivery) {

	// Establish a pool connection
	pool, err := pgxpool.New(context.Background(), db.GetConnString())
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer pool.Close()

	if err != nil {
		fmt.Printf("erro executando insert " + err.Error())
	} else {
		fmt.Printf("inserido com sucesso")
	}

	if err := execAssignStatements(context.Background(), pool, delivery); err != nil {
		log.Printf("Assign Delivery FAILED: %v\n", err)
	}

}
func execAssignStatements(ctx context.Context, pool *pgxpool.Pool, delivery model.Delivery) error {
	// 1. Begin the Transaction
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	// 2. Defer Rollback
	// Use a named return value or an outer variable to capture the commit error.
	defer func() {
		if r := recover(); r != nil {
			// Handle panic inside the transaction
			tx.Rollback(ctx)
			panic(r)
		} else if err != nil {
			// Rollback if an error occurred in the body
			tx.Rollback(ctx)
		} else {
			// If no error, attempt to commit
			if commitErr := tx.Commit(ctx); commitErr != nil {
				// Capture commit error
				err = fmt.Errorf("transaction commit failed: %w", commitErr)
			}
		}
	}()

	var deliveryAssignmentId uuid.UUID

	err = pool.QueryRow(
		ctx,
		`INSERT INTO delivery_assignments(driver_id, vehicle_id, status, route_description) 
		VALUES ($1, (select id from vehicles where plate_number = $2), $3, $4) RETURNING id`,
		delivery.DriverUUID, delivery.VehiclePlate, delivery.Status, delivery.RouteDescription,
	).Scan(&deliveryAssignmentId)

	if err != nil {
		return fmt.Errorf("Failed to insert on delivery_assignments %w", err)
	}

	_, err = tx.Exec(
		ctx,
		`UPDATE drivers SET status = 'IN_TRANSIT' WHERE id = $1`,
		delivery.DriverUUID,
	)
	if err != nil {
		return fmt.Errorf("Failed to update driver %w", err)
	}

	_, err = tx.Exec(
		ctx,
		`UPDATE vehicles SET status = 'IN_TRANSIT' WHERE plate_number = $1`,
		delivery.VehiclePlate,
	)
	if err != nil {
		return fmt.Errorf("Failed to update on vehicle %w", err)
	}

	for _, invoice := range delivery.Invoices {
		_, err = tx.Exec(
			ctx,
			`INSERT INTO delivery_invoices (delivery_assignment_id, invoice_number) VALUES ($1, $2)`,
			deliveryAssignmentId, invoice,
		)
		if err != nil {
			return fmt.Errorf("Failed to insert on delivery_invoices %w", err)
		}
	}

	return nil
}
func IsDriverAvailable(id string) (bool, error) {
	// Establish a pool connection
	pool, err := pgxpool.New(context.Background(), db.GetConnString())
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer pool.Close()

	sql := `SELECT status FROM drivers WHERE id = $1`

	// Execute the Query
	rows, err := pool.Query(context.Background(), sql, id)

	if err != nil {
		return false, fmt.Errorf("query failed: %w", err)
	}
	// IMPORTANT: Close the rows when done to release the connection back to the pool
	defer rows.Close()

	var driverStatus string
	// Iterate and scan results
	for rows.Next() {
		if err := rows.Scan(&driverStatus); err != nil {
			return false, fmt.Errorf("row scan failed: %w", err)
		}
	}

	// Check for any errors that occurred during row iteration
	if err := rows.Err(); err != nil {
		return false, fmt.Errorf("row iteration error: %w", err)
	}

	return driverStatus == "AVAILABLE", nil
}

// func isVehicleAvailable() bool {
// 	sql := `SELECT status FROM vehicles WHERE id = $1`

// 	// Execute the Query
// 	rows, err := pool.Query(ctx, sql, minID)
// }
