package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleCreateBinary - create binary data
func (h *Handler) HandleCreateBinary(ctx context.Context, req *grpc.CreateTextRequest) (*grpc.CreateTextResponse, error) {
	h.logger.Info("Create binary")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.CreateTextResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	return &grpc.CreateTextResponse{Text: text}, nil
}
