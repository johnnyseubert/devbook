package routes

import (
	"net/http"

	"github.com/johnnyseubert/devbook/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.GetAllUsers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.GetUserById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},
}
