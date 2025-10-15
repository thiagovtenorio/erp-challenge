package model

import (
	helper "api-erp-go/helper"
	"encoding/json"
	"log"
)

type Delivery struct {
	DriverUUID       string
	VehiclePlate     string
	Status           string
	RouteDescription string
	InvoicesJson     string
	Invoices         []string
}

func ParseJsonToStruct(jsonString string) Delivery {
	result := helper.JsonBodyToMap(jsonString)

	delivery := Delivery{
		DriverUUID:       result["driverId"].(string),
		VehiclePlate:     result["vehiclePlate"].(string),
		Status:           "ASSIGNED",
		RouteDescription: result["routeDescription"].(string),
		InvoicesJson:     result["invoices"].(string),
	}

	if err := json.Unmarshal([]byte(delivery.InvoicesJson), &delivery.Invoices); err != nil {
		log.Fatal(err)
	}

	return delivery
}
