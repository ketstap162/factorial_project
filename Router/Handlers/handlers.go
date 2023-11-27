package Handlers

import (
	"encoding/json"
	"factorial_project/Factorial"
	"factorial_project/HttpUtils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CalculateHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Parse JSON request body
	var requestBody HttpUtils.FactirualRequestBody
	decoder := json.NewDecoder(request.Body)
	error := decoder.Decode(&requestBody)

	if error != nil {
		http.Error(writer, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if requestBody.A < 0 || requestBody.B < 0 {
		http.Error(writer, `{"error":"Incorrect input"}`, http.StatusBadRequest)
		return
	}

	result := Factorial.CalculateFactorials(requestBody.A, requestBody.B)

	response := HttpUtils.FactorialResponseBody{
		AFactorial: result["a_factorial"],
		BFactorial: result["b_factorial"],
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
