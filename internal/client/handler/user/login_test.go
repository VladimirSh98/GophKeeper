package user

import (
	"testing"

	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user_service"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestLoginCmdRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockUser.NewMockServiceInterface(ctrl)
	handler := &Handler{
		userService: mockService,
	}

	cmd := handler.LoginCmd()

	mockService.EXPECT().Login(gomock.Any(), gomock.Any()).DoAndReturn(
		func(cmd *cobra.Command, args []string) {
			login, _ := cmd.Flags().GetString("login")
			require.Equal(t, "testuser", login)

			password, _ := cmd.Flags().GetString("password")
			require.Equal(t, "testpass", password)
		},
	)

	cmd.SetArgs([]string{"--login", "testuser", "--password", "testpass"})
	err := cmd.Execute()
	require.NoError(t, err)
}
