package routes

import (
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/pkg/user"
)

var usersRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       user.Create,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       user.SearchAll,
		Authentication: false,
	},
	{
		URI:            "/users/{userID}",
		Method:         http.MethodGet,
		Function:       user.Search,
		Authentication: false,
	},
	{
		URI:            "/users/{userID}",
		Method:         http.MethodPut,
		Function:       user.Update,
		Authentication: false,
	},
	{
		URI:            "/users/{userID}",
		Method:         http.MethodDelete,
		Function:       user.Delete,
		Authentication: false,
	},
}
