package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

func (c Event) EventUpload(name string, password string, file []byte, token model.Token) (string, error) {
	c.logger.Info("Upload binary data")

	//secretKey := encryption.AesKeySecureRandom([]byte(password))
	//encryptFile, err := encryption.Encrypt(string(file), secretKey)
	//if err != nil {
	//	c.logger.Error(err)
	//	return err, nil
	//}
	createdToken, _ := service.ConvertTimeToTimestamp(token.CreatedAt)
	endDateToken, _ := service.ConvertTimeToTimestamp(token.EndDateAt)
	uploadFile, err := c.grpc.HandleUploadBinary(context.Background(),
		&grpc.UploadBinaryRequest{Name: name, Data: []byte(file),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	c.logger.Debug(uploadFile.Name)
	return uploadFile.Name, nil
}