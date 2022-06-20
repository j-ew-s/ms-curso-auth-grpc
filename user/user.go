package user

import (
	"context"
)

type Server struct {
}

/*
Estes métodos tem que respeitar a interface que está no user.pb.go
Então pode copiar o contrato e colar aqui, se quiser, remove o opt.
Desta maneira, Server está obedecendo o contrato então ele passa a ser do tipo
necessário para se executar no nosso server.go
*/
func (s *Server) GetUser(ctx context.Context, in *UserId) (*User, error) {
	if in.Id == "1" {
		return &User{
			Id:    "1",
			Name:  "A",
			Email: "a@a.com",
			Address: &Address{
				Street:  "Rua A",
				ZipCode: "1234567",
			},
			Account: &Account{
				User:     "user",
				Password: "123",
			},
		}, nil
	}

	return new(User), nil
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
