package database

import (
	"fmt"

	authModels "github.com/j-ew-s/ms-curso-auth-grpc/auth-services/models"
)

func (sqlCommand SQLCommand) IsTokenValid(input string) (*authModels.Authorization, error) {

	authorization := authModels.Authorization{}

	db, err := sqlCommand.OpenConnection()
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		return &authorization, err
	}

	err = db.Table("login").
		Select("username, login_limit").
		Where("token = ? and login_limit > NOW()", input).
		Scan(&authorization).
		Error

	return &authorization, err
}
