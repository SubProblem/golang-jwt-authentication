package service

import (
	"subproblem/rest-api/pkg/db"
	"subproblem/rest-api/pkg/dto"
)

type UserService struct {
	db db.PostgresDb
}

func NewUserService(database *db.PostgresDb) *UserService {
	return &UserService{
		db: *database,
	}
}

func (u *UserService) GetUserById(id int) (*dto.UserResponseDto, error) {

	user, err := u.db.GetUserById(id)

	if err != nil {
		return nil, err
	}

	userDto := dto.ToUserDto(*user)

	return userDto, nil
}

func (u *UserService) AddUser(userDto dto.UserRequestDto) error {
	
	user := dto.ToUser(userDto)

	err := u.db.AddUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) DeleteUser(id int) (*dto.UserResponseDto, error) {

	user, err := u.db.DeleteUserById(id)

	if err != nil {
		return nil, err
	}

	userResponse := dto.ToUserDto(*user)

	return userResponse, nil
}