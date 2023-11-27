package main

import (
	"factorial_project/Router"
	"fmt"
	"net/http"
)

func main() {
	router := Router.GetRouter()

	// Start the server on port 8989
	fmt.Println("Server listening on :8989...")
	http.ListenAndServe(":8989", router)
}
