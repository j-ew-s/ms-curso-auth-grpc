package auth

import (
	"context"
	fmt "fmt"

	"github.com/j-ew-s/ms-curso-auth-grpc/database"
)

type AuthService struct {
	DBConn *database.SQLCommand
}

func (as *AuthService) IsTokenValid(ctx context.Context, in *Token) (*TokenValidation, error) {
	if in.Token == "" {
		return &TokenValidation{
			Status:  false,
			Message: "Token was empty",
		}, nil
	}

	authorization, err := as.DBConn.IsTokenValid(in.Token)

	if err != nil {
		fmt.Println(fmt.Printf(" IS TOKEN VALID ERROR :::  %+v ::", err))
		return &TokenValidation{
			Status:  false,
			Message: "Error when validating tokne",
		}, err
	}

	if len(authorization.Username) > 0 {

		return &TokenValidation{
			Status:  true,
			Message: fmt.Sprintf("Token is valid until : %s", fmt.Sprint(authorization.LoginLimit)),
		}, err
	}

	return &TokenValidation{
		Status:  false,
		Message: fmt.Sprintf("NOT FOUND TOKEN : %s", in.Token),
	}, err

}
