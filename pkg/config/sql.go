package config

import (
	"fmt"
	"log"
	"os"
)

var (
	dbUser                   = os.Getenv("DB_USER")
	dbPwd                    = os.Getenv("DB_PASS")
	dbInstanceConnectionName = os.Getenv("DB_INSTANCE_CONNECTION_NAME")
	dbName                   = os.Getenv("DB_NAME")
)

func GetDbUri() string {
	log.Printf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, dbInstanceConnectionName, dbName)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, dbInstanceConnectionName, dbName)
}
