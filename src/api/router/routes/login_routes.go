package routes

import (
	"net/http"
	"src/blogos/src/api/controllers"
)

var loginRoutes = []Route{
	{
		Uri:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
