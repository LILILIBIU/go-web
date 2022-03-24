package dto

import "Common/SQL"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user *SQL.User) *UserDto {
	return &UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
