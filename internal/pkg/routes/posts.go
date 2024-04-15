package routes

import (
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/pkg/post"
)

var PostRoutes = []Route{
	{
		URI:            "/posts",
		Method:         http.MethodPost,
		Function:       post.Create,
		Authentication: true,
	},
	{
		URI:            "/posts",
		Method:         http.MethodGet,
		Function:       post.SearchAll,
		Authentication: true,
	},
	{
		URI:            "/posts/{postID}",
		Method:         http.MethodGet,
		Function:       post.SearchOne,
		Authentication: true,
	},
	{
		URI:            "/posts/{postID}",
		Method:         http.MethodPut,
		Function:       post.Update,
		Authentication: true,
	},
	{
		URI:            "/posts/{postID}",
		Method:         http.MethodDelete,
		Function:       post.Delete,
		Authentication: true,
	},
}
