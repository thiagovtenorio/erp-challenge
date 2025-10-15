package helper

import (
	"encoding/json"
	"fmt"
	"log"
)

func JsonBodyToMap(jsonString string) map[string]interface{} {
	jsonData := []byte(jsonString)

	// Define a map to hold the dynamic data.
	var result map[string]interface{}

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON to map: %v", err)
	}

	fmt.Println("\nSuccessfully unmarshaled to map:")

	return result
}
