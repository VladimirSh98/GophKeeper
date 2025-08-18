package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunInitialization(t *testing.T) {
	app, err := Run()
	require.NoError(t, err)
	require.NotNil(t, app.RootCommand)
	require.NotNil(t, app.Logger)
	require.NotNil(t, app.Client)

	cmds := app.RootCommand.Commands()
	require.GreaterOrEqual(t, len(cmds), 9)

	names := make(map[string]bool)
	for _, c := range cmds {
		names[c.Use] = true
	}
	require.Contains(t, names, "register")
	require.Contains(t, names, "login")
	require.Contains(t, names, "delete user")
	require.Contains(t, names, "deleteSecret")
	require.Contains(t, names, "get secrets")
	require.Contains(t, names, "updateSecret")
	require.Contains(t, names, "createTextSecret")
	require.Contains(t, names, "createBinarySecret")
	require.Contains(t, names, "createLoginPassSecret")
	require.Contains(t, names, "createBankSecret")
}
