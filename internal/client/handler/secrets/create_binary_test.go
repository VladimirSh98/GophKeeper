package secrets

import (
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_service"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestCreateBinaryCmdRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockSecret.NewMockServiceInterface(ctrl)
	handler := &Handler{
		secretService: mockService,
	}

	cmd := handler.CreateBinaryCmd()

	mockService.EXPECT().CreateBinary(gomock.Any(), gomock.Any()).DoAndReturn(
		func(cmd *cobra.Command, args []string) {
			binary, _ := cmd.Flags().GetString("binary")
			require.Equal(t, "somebinarydata", binary)

			metadata, _ := cmd.Flags().GetStringArray("metadata")
			require.Equal(t, []string{"key=value"}, metadata)
		},
	)

	cmd.SetArgs([]string{
		"--binary", "somebinarydata",
		"--metadata", "key=value",
	})
	err := cmd.Execute()
	require.NoError(t, err)
}
