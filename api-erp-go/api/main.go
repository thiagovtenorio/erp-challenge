package main

import (
	"api-erp-go/model"
	assignmentRepo "api-erp-go/repository"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	// Register a handler function for the "/post" path
	http.HandleFunc("/v1/assignments/notify", handlePostRequest)

	fmt.Println("Server Running 8081 !")
	log.Fatal(http.ListenAndServe(":8081", nil))

}

// handlePostRequest is the handler function for POST requests to "/post"
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Print the received body (for demonstration purposes)
	fmt.Printf("Received POST request with body: %s\n", string(body))

	// Send a response back to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully received POST request with body: %s", string(body))

	delivery := model.ParseJsonToStruct(string(body))

	assignmentRepo.Insert(delivery)
}
