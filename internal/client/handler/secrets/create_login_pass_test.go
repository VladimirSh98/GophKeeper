package secrets

import (
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_service"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestCreateLoginPassCmdRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockSecret.NewMockServiceInterface(ctrl)
	handler := &Handler{
		secretService: mockService,
	}

	cmd := handler.CreateLoginPassCmd()

	mockService.EXPECT().CreateLoginPass(gomock.Any(), gomock.Any()).DoAndReturn(
		func(cmd *cobra.Command, args []string) {
			login, _ := cmd.Flags().GetString("login")
			require.Equal(t, "testlogin", login)

			password, _ := cmd.Flags().GetString("password")
			require.Equal(t, "testpass", password)

			metadata, _ := cmd.Flags().GetStringArray("metadata")
			require.Equal(t, []string{"role=admin"}, metadata)
		},
	)

	cmd.SetArgs([]string{
		"--login", "testlogin",
		"--password", "testpass",
		"--metadata", "role=admin",
	})
	err := cmd.Execute()
	require.NoError(t, err)
}
