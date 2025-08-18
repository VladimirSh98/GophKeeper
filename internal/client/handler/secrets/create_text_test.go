package secrets

import (
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_service"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestCreateTextCmdRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockSecret.NewMockServiceInterface(ctrl)
	handler := &Handler{
		secretService: mockService,
	}

	cmd := handler.CreateTextCmd()

	mockService.EXPECT().CreateText(gomock.Any(), gomock.Any()).DoAndReturn(
		func(cmd *cobra.Command, args []string) {
			text, _ := cmd.Flags().GetString("text")
			require.Equal(t, "test text", text)

			metadata, _ := cmd.Flags().GetStringArray("metadata")
			require.Equal(t, []string{"category=notes"}, metadata)
		},
	)

	cmd.SetArgs([]string{
		"--text", "test text",
		"--metadata", "category=notes",
	})
	err := cmd.Execute()
	require.NoError(t, err)
}
