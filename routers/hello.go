package routers

import (
	"todone-api/controllers"
	"todone-api/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"todone-api/core/authorization"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(authorization.CheckResourceAuthorization),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
