package router

import (
	"github.com/gorilla/mux"
	"github.com/johnnyseubert/devbook/src/router/routes"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigureRoutes(r)
}
