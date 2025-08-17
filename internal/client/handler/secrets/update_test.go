package secrets

import (
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_service"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestUpdateCmdRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockSecret.NewMockServiceInterface(ctrl)
	handler := &Handler{
		secretService: mockService,
	}

	cmd := handler.UpdateCmd()

	mockService.EXPECT().Update(gomock.Any(), gomock.Any()).DoAndReturn(
		func(cmd *cobra.Command, args []string) {
			secretID, _ := cmd.Flags().GetInt("secret")
			require.Equal(t, 42, secretID)

			number, _ := cmd.Flags().GetString("number")
			require.Equal(t, "1234567890123456", number)

			expiryMonth, _ := cmd.Flags().GetString("expiry_month")
			require.Equal(t, "12", expiryMonth)

			expiryYear, _ := cmd.Flags().GetString("expiry_year")
			require.Equal(t, "2025", expiryYear)

			holderName, _ := cmd.Flags().GetString("holder_name")
			require.Equal(t, "John Doe", holderName)

			cvv, _ := cmd.Flags().GetString("cvv")
			require.Equal(t, "123", cvv)

			login, _ := cmd.Flags().GetString("login")
			require.Equal(t, "user123", login)

			password, _ := cmd.Flags().GetString("password")
			require.Equal(t, "pass123", password)

			text, _ := cmd.Flags().GetString("text")
			require.Equal(t, "some text", text)

			binary, _ := cmd.Flags().GetString("binary")
			require.Equal(t, "binarydata", binary)
		},
	)

	cmd.SetArgs([]string{
		"--secret", "42",
		"--number", "1234567890123456",
		"--expiry_month", "12",
		"--expiry_year", "2025",
		"--holder_name", "John Doe",
		"--cvv", "123",
		"--login", "user123",
		"--password", "pass123",
		"--text", "some text",
		"--binary", "binarydata",
	})
	err := cmd.Execute()
	require.NoError(t, err)
}
