package routes

import (
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/pkg/userRunner"
)

var Users = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       userRunner.Create,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       userRunner.SearchAll,
		Authentication: true,
	},
	{
		URI:            "/users/{userID}",
		Method:         http.MethodGet,
		Function:       userRunner.Search,
		Authentication: true,
	},
	{
		URI:            "/users/{userID}",
		Method:         http.MethodPut,
		Function:       userRunner.Update,
		Authentication: true,
	},
	{
		URI:            "/users/{userID}",
		Method:         http.MethodDelete,
		Function:       userRunner.Delete,
		Authentication: true,
	},
	{
		URI:            "/users/{userID}/follow",
		Method:         http.MethodPost,
		Function:       userRunner.Follow,
		Authentication: true,
	},
	{
		URI:            "/users/{userID}/unfollow",
		Method:         http.MethodPost,
		Function:       userRunner.Unfollow,
		Authentication: true,
	},
}
