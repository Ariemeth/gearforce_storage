package config

import (
	"fmt"
	"os"
)

var envDBAddressName = "DB_ADDRESS"
var envDBUsername = "DB_USER"
var envDBPassword = "DB_PASS"
var envPort = "PORT"

var defaultPort = "9000"

type Main struct {
	Database Database
	System   System
}

type Database struct {
	Address  string
	User     string
	Password string
}

type System struct {
	Port string
}

func LoadEnvConfig() (Main, error) {

	config := Main{}

	DBConfig := Database{}

	val, ok := os.LookupEnv(envDBAddressName)
	if !ok {
		return config, fmt.Errorf("no database address environment variable [%s] found", envDBAddressName)
	}
	DBConfig.Address = val

	val, ok = os.LookupEnv(envDBUsername)
	if !ok {
		return config, fmt.Errorf("no database address environment variable [%s] found", envDBUsername)
	}
	DBConfig.User = val

	val, ok = os.LookupEnv(envDBPassword)
	if !ok {
		return config, fmt.Errorf("no database address environment variable [%s] found", envDBPassword)
	}
	DBConfig.Password = val

	config.Database = DBConfig

	SystemConfig := System{Port: defaultPort}
	val, ok = os.LookupEnv(envPort)
	if ok {
		SystemConfig.Port = val
	}

	config.System = SystemConfig

	return config, nil
}
