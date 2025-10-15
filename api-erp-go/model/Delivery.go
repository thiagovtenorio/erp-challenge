package model

import (
	helper "api-erp-go/helper"
)

type Delivery struct {
	DriverUUID       string
	VehiclePlate     string
	Status           string
	RouteDescription string
	Invoices         []string
}

func ParseJsonToStruct(jsonString string) Delivery {
	result := helper.JsonBodyToMap(jsonString)
	//result["invoices"]

	delivery := Delivery{
		DriverUUID:       result["driverId"].(string),
		VehiclePlate:     result["vehiclePlate"].(string),
		Status:           "ASSIGNED",
		RouteDescription: result["routeDescription"].(string),
	}

	return delivery
}
