package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johnnyseubert/devbook/src/middlewares"
)

type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func ConfigureRoutes(r *mux.Router) *mux.Router {
	var routes []Route

	routes = append(routes, loginRoute)
	routes = append(routes, userRoutes...)

	for _, route := range routes {

		if route.RequireAuthentication {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(
					middlewares.Authenticate(
						route.Function,
					),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(
					route.Function,
				),
			).Methods(route.Method)
		}
	}

	return r
}
