package repository

import (
	db "api-erp-go/db-connection"
	"api-erp-go/model"
	"context"
	"fmt"
)

// INSERT INTO delivery_assignments(driver_id, vehicle_id, status, route_description)
// VALUES (driver.id, vehicle.id, "ASSIGNED", "Entrega de produtos frágeis na região central") .
func Insert(delivery model.Delivery) {
	conn, _ := db.Connect()

	sqlStatement := `INSERT INTO delivery_assignments(driver_id, vehicle_id, status, route_description) 
		VALUES ($1, (select id from vehicles where plate_number = $2), $3, $4)`

	_, err := conn.Exec(context.Background(), sqlStatement, delivery.DriverUUID, delivery.VehiclePlate, delivery.Status, delivery.RouteDescription)

	if err != nil {
		fmt.Printf("erro executando insert " + err.Error())
	} else {
		fmt.Printf("inserido com sucesso")
	}

}

//UPDATE drivers SET status = "IN_TRANSIT" WHERE id = driver.id
//UPDATE vehicle SET status = "IN_TRANSIT" WHERE id = vehicle.id.
//INSERT INTO delivery_invoices (delivery_assignments_id, invoice) VALUES (deliveryAssignment.id, invoice)
