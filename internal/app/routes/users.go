package routes

import (
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/app/controllers"
)

var usersRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreatingUser,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.SearchUsers,
		Authentication: false,
	},
	{
		URI:            "/users/{userID}",
		Method:         http.MethodGet,
		Function:       controllers.SearchUser,
		Authentication: false,
	},
	{
		URI:            "/users/{userID}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		Authentication: false,
	},
	{
		URI:            "/users/{userID}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		Authentication: false,
	},
}
