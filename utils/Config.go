// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package utils

import (
	"fmt"
	"os"
	"strconv"
)

// Required application environment variables.
type ApplicationEnvironment struct {
	ApplicationID             string
	ApplicationName           string
	ApplicationTimezone       string
	GoScopeDatabaseConnection string
	GoScopeDatabaseType       string
	GoScopeEntriesPerPage     int
}

// Config is the global instance of the application's configuration.
var Config ApplicationEnvironment //nolint:gochecknoglobals

// Initialize the configuration instance to the values described in the .env file.
func ConfigSetup() {
	Config = ApplicationEnvironment{
		ApplicationID:             getString("APPLICATION_ID"),
		ApplicationName:           getString("APPLICATION_NAME"),
		ApplicationTimezone:       getString("APPLICATION_TIMEZONE"),
		GoScopeDatabaseConnection: getString("GOSCOPE_DATABASE_CONNECTION"),
		GoScopeDatabaseType:       getString("GOSCOPE_DATABASE_TYPE"),
		GoScopeEntriesPerPage:     getInteger("GOSCOPE_ENTRIES_PER_PAGE"),
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

	integer, _ := strconv.ParseInt(os.Getenv(key), 10, 32)

	return int(integer)
}
