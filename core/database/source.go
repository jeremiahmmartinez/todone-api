package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // This is for MySQL initialization.
	"todone-api/settings"
)

func GetDataSource () string {
	databaseSettings := settings.Get().Database

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", databaseSettings["user"], databaseSettings["password"],
		databaseSettings["host"], databaseSettings["port"], databaseSettings["name"])
}

type Adapter struct {
	driverName     string
	dataSourceName string
	DB             *sql.DB
}

// NewAdapter is the constructor for adapter.
func newDatabaseAdapter(driverName string, dataSourceName string) *Adapter {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName
	return a
}

var adapter *Adapter

func GetDatabaseAdapter() *Adapter {
	if adapter == nil {
		adapter = newDatabaseAdapter("mysql", GetDataSource())
	}

	return adapter
}

func (a *Adapter) Open() {
	db, err := sql.Open(a.driverName, a.dataSourceName)

	if err != nil {
		panic(err)
	}

	a.DB = db
}

func (a *Adapter) Close() {
	a.DB.Close()
}