package routers

import (
	"github.com/gorilla/mux"
	"fmt"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetUserRoutes(router)
	router = SetAuthenticationRoutes(router)
	fmt.Println("Here.")
	return router
}
