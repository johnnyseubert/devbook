package routes

import (
	"net/http"

	"github.com/johnnyseubert/devbook/src/controllers"
)

var loginRoute = Route{
	URI:                   "/login",
	Method:                http.MethodPost,
	Function:              controllers.Login,
	RequireAuthentication: false,
}
