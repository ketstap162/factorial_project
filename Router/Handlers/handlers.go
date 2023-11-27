package Handlers

import (
	"encoding/json"
	"factorial_project/Factorial"
	"factorial_project/HttpUtils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse JSON request body
	var requestBody HttpUtils.FactirualRequestBody
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

	result := Factorial.CalculateFactorials(requestBody.A, requestBody.B)

	response := HttpUtils.FactorialResponseBody{
		AFactorial: result["a_factorial"],
		BFactorial: result["b_factorial"],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
