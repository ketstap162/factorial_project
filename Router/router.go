package Router

import (
	"factorial_project/Router/Handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetRouter() *httprouter.Router {
	router := httprouter.New()

	router.POST("/calculate", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}
		Handlers.CalculateHandler(w, r, ps)
	})
	return router
}
