package routes

import (
	"car_rent/views"
	"github.com/gorilla/mux"
)

type Static struct {
	urlPath string
	filePath string
	layout string
}

const (
	defaultLayout = "bootstrap"
)

//Add more static pages here if needed
var (
	static = []Static {
		{
		urlPath:  "/",
		filePath: "static/home",
		layout:   defaultLayout,
		},
	}
)

func createStaticRoutes(r *mux.Router) {
	for _, conf := range static {
		tpl := views.NewView(conf.layout, conf.filePath)
		r.Handle(conf.urlPath, tpl)
	}
}
