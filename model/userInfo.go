package model

import (
	"time"
)

type UserInfo struct {
	UserId        string    `gorm:"primaryKey" json:"userId" `
	UserNumber    string    `json:"userNumber"`
	Password      string    `json:"password"`
	DeptName      string    `json:"deptName"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	SuperiorId    string    `json:"superiorId"`
	Type          int       `json:"type"`
	HeadUrl       string    `json:"headUrl"`
	Integral      int       `json:"integral"`
	ValidIntegral int       `json:"validIntegral"`
	Expiration    int       `json:"expiration"`
	CreateTime    time.Time `json:"createTime"`
	UpdateTime    time.Time `json:"updateTime"`
}

func QueryUser(userId string) (*UserInfo, error) {
	var user UserInfo
	DB.Model(UserInfo{}).First(&user, UserInfo{UserId: userId})
	return &user, nil
}
