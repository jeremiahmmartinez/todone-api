package policy

import (
	"github.com/jeremiahmmartinez/casbin"
	"github.com/jeremiahmmartinez/casbin-postgres-adapter"
	"todone-api/core/database"
)

var enforcer *casbin.Enforcer

func LoadPolicy() error {
	dbAdapter := adapter.NewFilteredAdapter("mysql", database.GetDataSource())

	enforcer = casbin.NewEnforcer()

	enforcer.InitWithAdapter("core/authorization/casbin/policy.conf", dbAdapter)

	err := enforcer.LoadPolicy()
	return err
}

func IsUserAuthorized (username string, resource string, action string) (bool, error) {
	if err := LoadPolicyByUser(username); err != nil {
		return false, err
	}

	return enforcer.Enforce(username, resource, action), nil
}

func LoadPolicyByUser(username string) error {
	dbAdapter := adapter.NewFilteredAdapter("mysql", database.GetDataSource())

	enforcer = casbin.NewEnforcer()

	enforcer.InitWithAdapter("core/authorization/casbin/policy.conf", dbAdapter)

	policyFilter := map[string]interface{}{"v0": username}

	err := enforcer.LoadFilteredPolicy(policyFilter)

	return err
}

func AddPolicyLine(rule []string) error {
	dbAdapter := adapter.NewFilteredAdapter("mysql", database.GetDataSource())

	err := dbAdapter.AddPolicy("p", "p", rule)

	return err
}

func RemovePolicyLine(rule []string) error {
	dbAdapter := adapter.NewFilteredAdapter("mysql", database.GetDataSource())

	err := dbAdapter.RemovePolicy("p", "p", rule)

	return err
}