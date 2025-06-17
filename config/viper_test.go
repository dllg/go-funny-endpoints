package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestDefaults(t *testing.T) {
	viper.Reset()
	Init()

	// Default ProjectID
	require.Equal(t, "funny-endpoints", Get(ProjectID))
	// No default set for these, so should be empty/zero/false
	require.Empty(t, Get(VersionBuildCommit))
	require.Equal(t, 0, GetInt(VersionBuildCommit))
	require.False(t, GetBool(VersionBuildRelease))
}

func TestEnvOverrideString(t *testing.T) {
	os.Setenv("SOME_PROJECT", "foo")
	defer os.Unsetenv("SOME_PROJECT")

	viper.Reset()
	Init()

	require.Equal(t, "foo", Get(ProjectID))
}

func TestEnvOverrideInt(t *testing.T) {
	os.Setenv("VERSION_BUILD_COMMIT", "12345")
	defer os.Unsetenv("VERSION_BUILD_COMMIT")

	viper.Reset()
	Init()

	require.Equal(t, "12345", Get(VersionBuildCommit))
	require.Equal(t, 12345, GetInt(VersionBuildCommit))
}

func TestEnvOverrideBool(t *testing.T) {
	os.Setenv("VERSION_BUILD_RELEASE", "true")
	defer os.Unsetenv("VERSION_BUILD_RELEASE")

	viper.Reset()
	Init()

	require.True(t, GetBool(VersionBuildRelease))
}

func TestServerPortDefault(t *testing.T) {
	// reset Viper and re‐apply defaults
	viper.Reset()
	Init()

	got := GetInt(ServerPort)
	wantDefault := 18080 // ← change this to your actual default

	if got != wantDefault {
		t.Errorf("default %q = %d; want %d", ServerPort, got, wantDefault)
	}
}

func TestServerPortEnvOverride(t *testing.T) {
	// set the env var that Viper will bind to "server.port"
	os.Setenv("SERVER_PORT", "9090")
	defer os.Unsetenv("SERVER_PORT")

	viper.Reset()
	Init()

	got := GetInt(ServerPort)
	want := 9090

	if got != want {
		t.Errorf("env override %q = %d; want %d", ServerPort, got, want)
	}
}
