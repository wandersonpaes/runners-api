package routes

import (
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/pkg/auth"
)

var Login = Route{
	URI:            "/login",
	Method:         http.MethodPost,
	Function:       auth.Login,
	Authentication: false,
}
