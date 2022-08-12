package service

import (
	"ginDemo/model"
)

type UserInfo struct {
	UserId string
}

func (user *UserInfo) QueryUser() (*model.UserInfo, error) {
	var result *model.UserInfo

	result, err := model.QueryUser(user.UserId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
