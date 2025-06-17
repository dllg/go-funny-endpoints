package config

import (
	"strings"

	"github.com/spf13/viper"
)

type js = map[string]interface{}

// Readers/writers will use these names for the configs to avoid hardcoding
const (
	ProjectID            = "some.project"
	VersionBuildCommit   = "version.build.commit"
	VersionBuildDateTime = "version.build.date.time"
	VersionBuildRelease  = "version.build.release"
	VersionGitHash       = "version.git.hash"
	ServerPort           = "server.port"
	URL                  = "url"
)

// Defaults provides sane default configurations when possible
var Defaults = js{
	ProjectID:  "funny-endpoints",
	ServerPort: 18080,
	URL:        "http://localhost:18080",
}

// Init will configure viper configs
func Init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	// 1. Apply default configurations
	for k, v := range Defaults {
		viper.SetDefault(k, v)
	}

	// 2. Set configs based on ENV variables e.g. SERVER_PORT=8111
	viper.AutomaticEnv()

	// 3. Read the config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		// If the config file is not found, we can ignore the error
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err) // or handle it as needed
		}
	}
}

// Get methods are here to avoid adding viper dependency elsewhere in the codebase.

// Get string from config using a key
func Get(key string) string {
	return viper.GetString(key)
}

// GetBool gets the config using a key
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetInt gets the config using a key
func GetInt(key string) int {
	return viper.GetInt(key)
}
