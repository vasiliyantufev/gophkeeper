package events

import (
	"context"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/variables"
)

func (c Event) CardDelete(card []string, token model.Token) error {
	c.logger.Info("card delete")

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

	deletedCardEntityID, err := c.grpc.EntityDelete(context.Background(),
		&grpc.DeleteEntityRequest{Name: card[0], Type: variables.Card.ToString(),
			AccessToken: &grpc.Token{Token: token.AccessToken, UserId: token.UserID,
				CreatedAt: createdToken, EndDateAt: endDateToken}})
	if err != nil {
		c.logger.Error(err)
		return err
	}

	c.logger.Debug(deletedCardEntityID)
	return nil
}
