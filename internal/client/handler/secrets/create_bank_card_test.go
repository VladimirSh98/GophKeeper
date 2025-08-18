package secrets

import (
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_service"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestCreateBankCardCmdRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockSecret.NewMockServiceInterface(ctrl)
	handler := &Handler{
		secretService: mockService,
	}

	cmd := handler.CreateBankCardCmd()

	mockService.EXPECT().CreateBankCard(gomock.Any(), gomock.Any()).DoAndReturn(
		func(cmd *cobra.Command, args []string) {
			number, _ := cmd.Flags().GetString("number")
			require.Equal(t, "1234567890123456", number)

			expiryMonth, _ := cmd.Flags().GetString("expiry_month")
			require.Equal(t, "12", expiryMonth)

			expiryYear, _ := cmd.Flags().GetString("expiry_year")
			require.Equal(t, "2030", expiryYear)

			holderName, _ := cmd.Flags().GetString("holder_name")
			require.Equal(t, "John Doe", holderName)

			cvv, _ := cmd.Flags().GetString("cvv")
			require.Equal(t, "123", cvv)

			metadata, _ := cmd.Flags().GetStringArray("metadata")
			require.Equal(t, []string{"key=value"}, metadata)
		},
	)

	cmd.SetArgs([]string{
		"--number", "1234567890123456",
		"--expiry_month", "12",
		"--expiry_year", "2030",
		"--holder_name", "John Doe",
		"--cvv", "123",
		"--metadata", "key=value",
	})
	err := cmd.Execute()
	require.NoError(t, err)
}
