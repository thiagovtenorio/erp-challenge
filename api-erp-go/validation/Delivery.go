package validation

import (
	"api-erp-go/model"
	assignment "api-erp-go/repository"
	"fmt"
)

/**Valida:
Se o driverId recebido existe e seu status é AVAILABLE.
Se o vehiclePlate recebido existe e o status do veículo é OPERATIONAL.

Se no campo invoice as notas fiscais possuem 44 dígitos e o dígito 21 a 22 é igual a 55.

Caso qualquer alguma validação falhe, retornar
       Status 422 — Unprocessable Entity com uma mensagem de erro clara no corpo da
resposta.*/

func IsValid(delivery model.Delivery) (bool, error) {

	driverAvailable, _ := assignment.IsDriverAvailable(delivery.DriverUUID)

	if !driverAvailable {
		return false, fmt.Errorf("unprocessable Entity: Driver not Available")
	} else {
		return true, nil
	}

}
