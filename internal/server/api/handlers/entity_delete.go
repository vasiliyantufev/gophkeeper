package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EntityDelete - check the validity of the token and delete record (text, bank card or login password)
func (h *Handler) EntityDelete(ctx context.Context, req *grpc.DeleteEntityRequest) (*grpc.DeleteEntityResponse, error) {
	h.logger.Info("delete entity")

	valid := h.token.Validate(req.AccessToken)
	if !valid {
		h.logger.Error(errors.ErrNotValidateToken)
		return &grpc.DeleteEntityResponse{}, status.Errorf(
			codes.Unauthenticated, errors.ErrNotValidateToken.Error(),
		)
	}

	DeletedEntityID, err := h.entity.Delete(req.AccessToken.UserId, req.Name, req.Type)
	if err != nil {
		h.logger.Error(err)
		return &grpc.DeleteEntityResponse{}, status.Errorf(
			codes.Internal, err.Error(),
		)
	}

	h.logger.Debug(DeletedEntityID)
	return &grpc.DeleteEntityResponse{Id: DeletedEntityID}, nil
}
