package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RequestBody is a struct representing the expected JSON structure in the request body.
type RequestBody struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ResponseBody is a struct representing the JSON response structure.
type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	// Define the HTTP handler function for handling POST requests.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST.
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Decode the JSON data from the request body into the RequestBody struct.
		var requestBody RequestBody
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestBody); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// Check if the expected "message" field is present in the JSON data.
		if requestBody.UserName == "" || requestBody.Email == "" || requestBody.Password == "" {
			http.Error(w, "Invalid JSON message", http.StatusBadRequest)
			return
		}

		// Print the received message to the server console.
		fmt.Printf("Received new user: %s\n", requestBody.UserName)
		username := requestBody.UserName
		password := requestBody.Password
		// Send a JSON response back to the client.
		response := ResponseBody{
			Status:  "success",
			Message: "User " + username + " successfully registered with password - " + password,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	// Start the HTTP server on port 8080.
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
