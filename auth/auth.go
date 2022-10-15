package auth

import (
	"context"
)

type Server struct {
}

func (s *Server) IsTokenValid(ctx context.Context, in *Token) (*TokenValidation, error) {
	if in.Token == "123" {
		return &TokenValidation{
			Status:  true,
			Message: "Valid",
		}, nil
	}

	return &TokenValidation{
		Status:  false,
		Message: "Invalid",
	}, nil
}
