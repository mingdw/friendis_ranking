package models

import (
	"fmt"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Age        int    `json:"age"`
	Sex        int    `json:"sex"`
	Phone      string `json:"phone"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"CreateTime"`
	UpEditTime int64  `json:"UpEditTime"`
	Creator    string `json:"Creator"`
	UpEditor   string `json:"UpEditor"`
}

func (User) TableName() string {
	return "sys_user"
}

func GetUserInfoById(id int) User {
	return initUser(id)
}

func initUser(id int) User {
	user := User{
		Id:         id,
		Username:   "张三",
		Password:   "pwdzhangsan",
		AddTime:    time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		Age:        20,
		Sex:        0,
		Phone:      "18710127234",
	}
	fmt.Println("user info: ", user)
	return user
}
