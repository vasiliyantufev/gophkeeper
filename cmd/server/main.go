package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/vasiliyantufev/gophkeeper/internal/server/api"
	grpcHandler "github.com/vasiliyantufev/gophkeeper/internal/server/api/handlers"
	"github.com/vasiliyantufev/gophkeeper/internal/server/config"
	"github.com/vasiliyantufev/gophkeeper/internal/server/database"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/card"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/metadata"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/text"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/token"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/repositories/user"
)

func main() {
	logger := logrus.New()
	config := config.NewConfig(logger)
	logger.SetLevel(config.DebugLevel)

	db, err := database.New(config, logger)
	if err != nil {
		logger.Fatal(err)
	} else {
		defer db.Close()
	}

	userRepository := user.New(db)
	textRepository := text.New(db)
	metadataRepository := metadata.New(db)
	cardRepository := card.New(db)
	tokenRepository := token.New(db)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	handlerGrpc := grpcHandler.NewHandler(db, userRepository, textRepository, cardRepository, metadataRepository, tokenRepository, logger)
	go api.StartService(handlerGrpc, config, logger)

	<-ctx.Done()
	logger.Info("server shutdown on signal with:", ctx.Err())
}
