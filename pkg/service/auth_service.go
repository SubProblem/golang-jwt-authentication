package service

import (
	"errors"
	"fmt"
	"subproblem/rest-api/pkg/db"
	"subproblem/rest-api/pkg/dto"
	"subproblem/rest-api/pkg/jwt"
	"subproblem/rest-api/pkg/util"
)

type AuthService struct{
	db db.PostgresDb
}

func NewAuthService(database *db.PostgresDb) *AuthService {
	return &AuthService{
		db: *database,
	}
}

func (auth *AuthService) Login(userReqeust dto.LoginRequestDto) (string, error) {

	
	// hashedPassword, _ := util.HashPassword(userReqeust.Password)
	
	user, err := auth.db.FindUserByEmail(userReqeust.Email)

	if err != nil {
		return "", err
	}
	

	if !util.ComparePassowrd(user.Password, userReqeust.Password) || user.Email != userReqeust.Email {
		return "", errors.New("Bad Credentials")
	}

	token, err := jwt.GenerateToken(userReqeust.Email)

	if err != nil {
		return "", nil
	}

	return token, nil
}


func (auth *AuthService) Register(userRequest dto.UserRequestDto) error {

	user, err := auth.db.FindUserByEmail(userRequest.Email)

	if err != nil {
		return err
	}
	fmt.Printf("user: %v\n", user)

	if user != nil{
		return errors.New("Registration Failed")
	}

	newUser := dto.ToUser(userRequest)

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = hashedPassword

	if err := auth.db.AddUser(newUser); err != nil {
		return errors.New("Registration Failed, try again")
	}

	return nil
}
