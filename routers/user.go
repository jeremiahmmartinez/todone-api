package routers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"todone-api/controllers"
	"todone-api/core/validation"
	"todone-api/core/authentication"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.Handle("/create-user",
		negroni.New(
			negroni.HandlerFunc(validation.ValidateNewUsernameFormat),
			negroni.HandlerFunc(validation.ValidateNewEmailFormat),
			negroni.HandlerFunc(validation.ValidateNewUsernameExists),
			negroni.HandlerFunc(validation.ValidateNewEmailExists),
			negroni.HandlerFunc(controllers.CreateUser),
		)).Methods("POST")

	router.Handle("/user/{userId}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			//negroni.HandlerFunc(validation.CheckUserRepoAuthorization),
			negroni.HandlerFunc(controllers.GetUser),
		)).Methods("GET")

	router.Handle("/users",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			//negroni.HandlerFunc(validation.CheckUserRepoAuthorization),
			negroni.HandlerFunc(controllers.GetUsers),
		)).Methods("GET")

	return router
}

