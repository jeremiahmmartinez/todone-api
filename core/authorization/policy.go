package authorization

import (
	"github.com/casbin/casbin"
	"github.com/jeremiahmmartinez/casbin-postgres-adapter"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Cloclo4694"
	dbname   = "casbin"
)

func loadPolicy() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	enforcer := casbin.NewEnforcer()

	dbAdapter := adapter.NewFilteredAdapter("postgres", psqlInfo)
	enforcer.InitWithAdapter("core/authorization/casbin/policy.conf", dbAdapter)


	error1 := enforcer.LoadPolicy()

	fmt.Println("LOAD POLICY")
	fmt.Println(error1)

}