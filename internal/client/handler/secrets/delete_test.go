package secrets

import (
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_service"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestDeleteCmdRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockSecret.NewMockServiceInterface(ctrl)
	handler := &Handler{
		secretService: mockService,
	}

	cmd := handler.DeleteCmd()

	mockService.EXPECT().Delete(gomock.Any(), gomock.Any()).DoAndReturn(
		func(cmd *cobra.Command, args []string) {
			secretID, _ := cmd.Flags().GetInt("secret")
			require.Equal(t, 42, secretID)
		},
	)

	cmd.SetArgs([]string{"--secret", "42"})
	err := cmd.Execute()
	require.NoError(t, err)
}
