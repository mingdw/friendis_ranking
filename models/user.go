package models

import (
	"fmt"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int64  `json:"addTime"`
	UpdateTime int64  `json:"updateTime"`
	Age        int    `json:"age"`
	Sex        int    `json:"sex"`
	Phone      string `json:"phone"`
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
