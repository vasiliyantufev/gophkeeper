package handlers

import (
	"context"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/proto"
)

// HandleGetListText - get list text
func (h *Handler) HandleGetListText(ctx context.Context, req *grpc.GetListTextRequest) (*grpc.GetListTextResponse, error) {
	var msg string

	ListText, err := h.text.GetListText(req.UserId)
	if err != nil {
		return &grpc.GetListTextResponse{}, err
	}
	h.logger.Info(ListText)

	return &grpc.GetListTextResponse{Message: msg}, nil
}
