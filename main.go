package main

import (
	"github.com/uadmin/uadmin"
)

func main() {

	// Initialize the database (and encryption)
	//! Do not remove this line
	Init()

	// Start the default uadmin server (at port 8080)
	uadmin.StartServer()
}
