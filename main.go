package main

import (
	"os"
	"strconv"

	"github.com/uadmin/uadmin"
)

func main() {

	// Map environment variables to the database config

	dbtype := "mysql"
	host := os.Getenv("MYSQLHOST")
	port, err := strconv.Atoi(os.Getenv("MYSQLPORT"))
	if err != nil {
		// ... handle error
		panic(err)
	}
	user := os.Getenv("MYSQLUSER")
	password := os.Getenv("MYSQLPASSWORD")
	name := os.Getenv("MYSQLDATABASE")

	// Connect to the database
	uadmin.Database = &uadmin.DBSettings{
		Type:     dbtype,
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
	}

	uadmin.StartServer()
}
