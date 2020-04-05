package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func addAssetsFileServer(r *mux.Router) {
	// Assets
	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)
}
