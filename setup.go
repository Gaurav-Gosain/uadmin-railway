package main

import (
	"encoding/json"
	"fmt"
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

	var key string
	var salt string

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

		fmt.Println("Encryption key and salt loaded from .encrypt file")

		key = encryptionSettings.Key
		salt = encryptionSettings.Salt

		defer encryptionConfigFile.Close()

	} else {

		fmt.Println("No .encrypt file found. Getting encryption key and salt from environment variables")

		key = os.Getenv("ENCRYPTIONKEY")
		salt = os.Getenv("SALT")

		uadmin.EncryptKey = []byte(os.Getenv("KEY"))
		uadmin.Salt = os.Getenv("SALT")
	}

	// Create a .salt file
	saltFile, err := os.Create(".salt")
	if err != nil {
		panic(err)
	}
	// Write the salt to the file
	_, err = saltFile.WriteString(salt)
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing salt to .salt file")

	defer saltFile.Close()

	// Create a .key file
	keyFile, err := os.Create(".key")
	if err != nil {
		panic(err)
	}
	// Write the key to the file as bytes
	_, err = keyFile.Write([]byte(key))
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing key to .key file")

	defer keyFile.Close()

	uadmin.EncryptKey = []byte(key)
	uadmin.Salt = salt
}

func Init() {
	// Initialize the database
	uadmin.Database = InitDatabase()

	//! Temporary workaround for:
	// [  ERROR ]   Hanlder.NewLogger. Unix syslog delivery error
	uadmin.LogTrail = false
	uadmin.LogHTTPRequests = false
	uadmin.LogRead = false
	uadmin.LogAdd = false
	uadmin.LogDelete = false
	uadmin.LogEdit = false
	uadmin.TrailLoggingLevel = uadmin.EMERGENCY

	// Set the encryption key and salt
	Encrypt()
}
