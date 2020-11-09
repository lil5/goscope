package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Required application environment variables.
type ApplicationEnvironment struct {
	ApplicationID                     string
	ApplicationName                   string
	ApplicationTimezone               string
	GoScopeDatabaseConnection         string
	GoScopeDatabaseType               string
	GoScopeEntriesPerPage             int
	HasFrontendDisabled               bool
	GoScopeDatabaseMaxOpenConnections int
	GoScopeDatabaseMaxIdleConnections int
	GoScopeDatabaseMaxConnLifetime    int
}

// Config is the global instance of the application's configuration.
var Config ApplicationEnvironment //nolint:gochecknoglobals

// Initialize the configuration instance to the values described in the .env file.
func ConfigSetup() {
	Config = ApplicationEnvironment{
		ApplicationID:                     getString("APPLICATION_ID"),
		ApplicationName:                   getString("APPLICATION_NAME"),
		ApplicationTimezone:               getString("APPLICATION_TIMEZONE"),
		GoScopeDatabaseConnection:         getString("GOSCOPE_DATABASE_CONNECTION"),
		GoScopeDatabaseType:               getString("GOSCOPE_DATABASE_TYPE"),
		GoScopeEntriesPerPage:             getInteger("GOSCOPE_ENTRIES_PER_PAGE"),
		HasFrontendDisabled:               getOptionalBool("GOSCOPE_DISABLE_FRONTEND", false),
		GoScopeDatabaseMaxOpenConnections: getOptionalInteger("GOSCOPE_DATABASE_MAX_OPEN_CONNECTIONS", 10),
		GoScopeDatabaseMaxIdleConnections: getOptionalInteger("GOSCOPE_DATABASE_MAX_IDLE_CONNECTIONS", 5),
		GoScopeDatabaseMaxConnLifetime:    getOptionalInteger("GOSCOPE_DATABASE_MAX_CONN_LIFETIME", 10),
	}
}

// Search for and return a string variable from the system environment, panicking if not present.
func getString(key string) string {
	_, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Error! Could not find %s environment variable!", key))
	}

	return os.Getenv(key)
}

// Search for and return an integer variable from the system environment, panicking if not present.
func getInteger(key string) int {
	_, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Error! Could not find %s environment variable!", key))
	}

	integer, err := strconv.ParseInt(os.Getenv(key), 10, 32)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s was not a valid integer!", key))
	}

	return int(integer)
}

// Search for and return an integer variable from the system environment, panicking if not present.
func getOptionalInteger(key string, defaultValue int) int {
	_, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	integer, err := strconv.ParseInt(os.Getenv(key), 10, 32)
	if err != nil {
		log.Printf("Error while reading integer %s from environment!", key)
		return defaultValue
	}

	return int(integer)
}

func getOptionalBool(key string, defaultValue bool) bool {
	_, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	b, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		log.Printf("Error while reading boolean %s from environment!", key)
		return defaultValue
	}

	return b
}
