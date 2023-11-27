package main

import (
	"encoding/json"
	"factorial_project/factorial"
	"fmt"
	"math/big"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RequestBody struct {
	A int `json:"a"`
	B int `json:"b"`
}

type ResponseBody struct {
	AFactorial *big.Int `json:"a_factorial"`
	BFactorial *big.Int `json:"b_factorial"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse JSON request body
	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if requestBody.A < 0 || requestBody.B < 0 {
		http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
		return
	}

	result := factorial.СalculateFactorials(requestBody.A, requestBody.B)

	// Close the channel once all goroutines are done

	// Prepare and send the response
	response := ResponseBody{
		AFactorial: result["a_factorial"],
		BFactorial: result["b_factorial"],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := httprouter.New()

	// Middleware to check if 'a' and 'b' exist in JSON and are non-negative integers
	router.POST("/calculate", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}
		calculateHandler(w, r, ps)
	})

	// Start the server on port 8989
	fmt.Println("Server listening on :8989...")
	http.ListenAndServe(":8989", router)
}
