package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	createStaticRoutes(router)
	addAssetsFileServer(router)

	return router
}
