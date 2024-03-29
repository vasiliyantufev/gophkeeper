package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/service/encryption"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

func (c Event) TextUpdate(name, passwordSecure, text string, token model.Token) error {
	c.logger.Info("text update")

	secretKey := encryption.AesKeySecureRandom([]byte(passwordSecure))
	encryptText, err := encryption.Encrypt(text, secretKey)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	createdToken, err := service.ConvertTimeToTimestamp(token.CreatedAt)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	endDateToken, err := service.ConvertTimeToTimestamp(token.EndDateAt)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	updatedTextEntityID, err := c.grpc.EntityUpdate(context.Background(),
		&grpc.UpdateEntityRequest{Name: name, Data: []byte(encryptText), Type: variables.Text.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID, CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(updatedTextEntityID)
	return nil
}
