package secret

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/client/utils"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"strings"
)

// Update update secret
func (s *Service) Update(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	token, err := s.tokenManager.GetToken()
	if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	if token == "" {
		utils.ColorMessage("Для удаления аккаунта необходимо авторизоваться", color.FgHiYellow)
		return
	}
	secretID, _ := cmd.Flags().GetInt("secret")
	var secret *pb.SecretModel
	secret, err = s.secretClient.GetByID(ctx, token, secretID)
	if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	content := getContentBySecret(cmd, secret)
	metadata := getMetadataBySecret(cmd, secret)
	_, err = s.secretClient.Update(ctx, token, secretID, content, metadata)
	if err != nil {
		utils.ColorMessage("Произошла непредвиденная ошибка", color.FgHiRed)
		return
	}
	utils.ColorMessage("Секрет обновлен", color.FgHiGreen)
}

func getContentBySecret(cmd *cobra.Command, secret *pb.SecretModel) []byte {
	contentData, err := parseSecret(secret.DataType, secret.Content)
	if err != nil {
		return secret.Content
	}
	switch secret.DataType {
	case pb.DataType_LOGIN_PASSWORD:
		return prepareLoginPass(cmd, contentData, secret.Content)
	case pb.DataType_BANK_CARD:
		return prepareBankCard(cmd, contentData, secret.Content)
	case pb.DataType_TEXT_DATA:
		return prepareText(cmd, contentData, secret.Content)
	case pb.DataType_BINARY_DATA:
		return prepareBinary(cmd, contentData, secret.Content)
	default:
		return secret.Content
	}

}

func prepareLoginPass(cmd *cobra.Command, contentData proto.Message, oldContent []byte) []byte {
	loginPass, ok := contentData.(*pb.LoginPass)
	if !ok {
		return oldContent
	}
	login, _ := cmd.Flags().GetString("login")
	password, _ := cmd.Flags().GetString("password")
	if login == "" {
		login = loginPass.Login
	}
	if password == "" {
		password = loginPass.Password
	}
	content, err := proto.Marshal(&pb.LoginPass{Login: login, Password: password})
	if err != nil {
		return oldContent
	}
	return content
}

func prepareBankCard(cmd *cobra.Command, contentData proto.Message, oldContent []byte) []byte {
	bankCard, ok := contentData.(*pb.BankCard)
	if !ok {
		return oldContent
	}
	number, _ := cmd.Flags().GetString("number")
	expiryMonth, _ := cmd.Flags().GetString("expiry_month")
	expiryYear, _ := cmd.Flags().GetString("expiry_year")
	holderName, _ := cmd.Flags().GetString("holder_name")
	cvv, _ := cmd.Flags().GetString("cvv")
	if number == "" {
		number = bankCard.Number
	}
	if expiryMonth == "" {
		expiryMonth = bankCard.ExpiryMonth
	}
	if expiryYear == "" {
		expiryYear = bankCard.ExpiryYear
	}
	if holderName == "" {
		holderName = bankCard.HolderName
	}
	if cvv == "" {
		cvv = bankCard.Cvv
	}
	content, err := proto.Marshal(&pb.BankCard{
		Number:      number,
		HolderName:  holderName,
		ExpiryMonth: expiryMonth,
		ExpiryYear:  expiryYear,
		Cvv:         cvv,
	})
	if err != nil {
		return oldContent
	}
	return content
}

func prepareText(cmd *cobra.Command, contentData proto.Message, oldContent []byte) []byte {
	textData, ok := contentData.(*pb.TextData)
	if !ok {
		return oldContent
	}
	text, _ := cmd.Flags().GetString("text")
	if text == "" {
		text = textData.Text
	}
	content, err := proto.Marshal(&pb.TextData{Text: text})
	if err != nil {
		return oldContent
	}
	return content
}

func prepareBinary(cmd *cobra.Command, contentData proto.Message, oldContent []byte) []byte {
	binaryData, ok := contentData.(*pb.BinaryData)
	if !ok {
		return oldContent
	}
	binaryString, _ := cmd.Flags().GetString("binary")
	binary := []byte(binaryString)
	if len(binary) == 0 {
		binary = binaryData.Data
	}
	content, err := proto.Marshal(&pb.BinaryData{Data: binary})
	if err != nil {
		return oldContent
	}
	return content
}

func getMetadataBySecret(cmd *cobra.Command, secret *pb.SecretModel) map[string]string {
	metadataArgs, _ := cmd.Flags().GetStringArray("metadata")
	metadata := make(map[string]string)
	for _, kv := range metadataArgs {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) == 2 {
			metadata[parts[0]] = parts[1]
		}
	}
	if len(metadata) == 0 {
		metadata = secret.Metadata
	}
	return metadata
}
