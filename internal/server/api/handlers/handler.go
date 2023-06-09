package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/card"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/metadata"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/text"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/token"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/user"
)

type Handler struct {
	database *database.DB
	user     *user.User
	text     *text.Text
	card     *card.Card
	metadata *metadata.Metadata
	token    *token.Token
	logger   *logrus.Logger
	grpc.UnimplementedGophkeeperServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(db *database.DB, userRepository *user.User, textRepository *text.Text, cardRepository *card.Card,
	metadataRepository *metadata.Metadata, tokenRepository *token.Token, log *logrus.Logger) *Handler {
	return &Handler{database: db, user: userRepository, text: textRepository, metadata: metadataRepository,
		card: cardRepository, token: tokenRepository, logger: log}
}
