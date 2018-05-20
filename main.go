package main

import (
	"todone-api/routers"
	"todone-api/settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

//Test GOOOOOO

func main() {
	settings.Init()

	/*User, err := user.GetUser(map[string]interface{}{"id": 1})

	fmt.Println(User)
	fmt.Println(err)
	jsonUser, err:= json.Marshal(User)
	fmt.Println(string(jsonUser))
	fmt.Println(err)*/


	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
