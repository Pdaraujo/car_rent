package routes

import (
	"car_rent/views"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		home := views.NewView("bootstrap", "static/home")
		home.Render(w, r, nil)
	})
	return router
}
