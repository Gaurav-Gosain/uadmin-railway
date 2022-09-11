package main

import (
	"encoding/json"
	"io"
	"os"
	"strconv"

	"github.com/uadmin/uadmin"
)

func InitDatabase() *uadmin.DBSettings {

	// Check if a local .database file exists
	if _, err := os.Stat(".database"); err == nil {
		// If it does, read the file and return the database settings

		// Step 1: Open the file
		databaseConfigFile, err := os.Open(".database")
		if err != nil {
			panic(err)
		}

		byteValue, _ := io.ReadAll(databaseConfigFile)

		// Step 2: Unmarshal the file
		var DBSettings uadmin.DBSettings

		err = json.Unmarshal(byteValue, &DBSettings)

		if err != nil {
			panic(err)
		}
		defer databaseConfigFile.Close()

		return &DBSettings
	}

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
	return &uadmin.DBSettings{
		Type:     dbtype,
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
	}
}

type EncryptionSettings struct {
	Key  string `json:"KEY"`
	Salt string `json:"SALT"`
}

func Encrypt() {
	// Check if a local .encrypt file exists
	if _, err := os.Stat(".encrypt"); err == nil {
		// If it does, read the file and return the database settings

		// Step 1: Open the file
		encryptionConfigFile, err := os.Open(".encrypt")
		if err != nil {
			panic(err)
		}

		byteValue, _ := io.ReadAll(encryptionConfigFile)

		// Step 2: Unmarshal the file
		var encryptionSettings EncryptionSettings

		err = json.Unmarshal(byteValue, &encryptionSettings)

		if err != nil {
			panic(err)
		}

		uadmin.EncryptKey = []byte(encryptionSettings.Key)
		uadmin.Salt = encryptionSettings.Salt

		defer encryptionConfigFile.Close()

		return
	}

	uadmin.EncryptKey = []byte(os.Getenv("KEY"))
	uadmin.Salt = os.Getenv("SALT")
}

func main() {

	// Initialize the database
	uadmin.Database = InitDatabase()
	// Set the encryption key and salt
	Encrypt()

	// Start the default uadmin server (at port 8080)
	uadmin.StartServer()
}
